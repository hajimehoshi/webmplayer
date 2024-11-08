// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 Hajime Hoshi

package main

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/ebml-go/webm"
)

type Stream struct {
	meta        webm.WebM
	videoStream *VideoStream
	audioStream *AudioStream

	reader *webm.Reader
}

func NewStream(r io.ReadSeeker) (*Stream, error) {
	s := &Stream{}
	reader, err := webm.Parse(r, &s.meta)
	if err != nil {
		return nil, err
	}
	s.reader = reader

	vTrack := s.meta.FindFirstVideoTrack()
	aTrack := s.meta.FindFirstAudioTrack()

	var vPackets chan webm.Packet
	var aPackets chan webm.Packet

	if vTrack != nil {
		vPackets = make(chan webm.Packet, 32)

		slog.Info(fmt.Sprintf("Found video track: %dx%d dur: %v %s", vTrack.DisplayWidth, vTrack.DisplayHeight, s.meta.Segment.GetDuration(), vTrack.CodecID))

		s.videoStream, err = NewVideoStream(VCodec(vTrack.CodecID), vPackets)
		if err != nil {
			return nil, err
		}
	}

	if aTrack != nil {
		aPackets = make(chan webm.Packet, 32)

		slog.Info(fmt.Sprintf("Found audio track: ch: %d %.1fHz, dur: %v, codec: %s", aTrack.Channels, aTrack.SamplingFrequency, s.meta.Segment.GetDuration(), aTrack.CodecID))

		s.audioStream, err = NewAudioDecoder(AudioCodec(aTrack.CodecID), aTrack.CodecPrivate,
			int(aTrack.Channels), int(aTrack.SamplingFrequency), aPackets)
		if err != nil {
			return nil, err
		}
	}

	go func() {
		for pkt := range s.reader.Chan {
			switch {
			case vTrack == nil:
				// Audio only.
				aPackets <- pkt
			case aTrack == nil:
				// Video Only.
				vPackets <- pkt
			default:
				switch pkt.TrackNumber {
				case vTrack.TrackNumber:
					vPackets <- pkt
				case aTrack.TrackNumber:
					aPackets <- pkt
				}
			}
		}
		close(vPackets)
		close(aPackets)
		s.reader.Shutdown()
	}()

	return s, nil
}

func (s *Stream) Meta() *webm.WebM {
	return &s.meta
}

func (s *Stream) VideoStream() *VideoStream {
	return s.videoStream
}

func (s *Stream) AudioStream() *AudioStream {
	return s.audioStream
}
