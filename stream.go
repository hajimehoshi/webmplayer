// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 Hajime Hoshi

package webmplayer

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/ebml-go/webm"
)

type stream struct {
	meta        webm.WebM
	videoStream *videoStream
	audioStream *audioStream

	reader *webm.Reader
}

func newStream(r io.ReadSeeker) (*stream, error) {
	s := &stream{}
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

		s.videoStream, err = newVideoStream(videoCodec(vTrack.CodecID), vPackets)
		if err != nil {
			return nil, err
		}
	}

	if aTrack != nil {
		aPackets = make(chan webm.Packet, 32)

		slog.Info(fmt.Sprintf("Found audio track: ch: %d %.1fHz, dur: %v, codec: %s", aTrack.Channels, aTrack.SamplingFrequency, s.meta.Segment.GetDuration(), aTrack.CodecID))

		s.audioStream, err = newAudioDecoder(audioCodec(aTrack.CodecID), aTrack.CodecPrivate,
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

func (s *stream) Meta() *webm.WebM {
	return &s.meta
}

func (s *stream) VideoStream() *videoStream {
	return s.videoStream
}

func (s *stream) AudioStream() *audioStream {
	return s.audioStream
}
