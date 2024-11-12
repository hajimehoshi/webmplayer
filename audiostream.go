// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 Hajime Hoshi

package webmplayer

import (
	"fmt"
	"unsafe"

	"github.com/ebml-go/webm"
	"github.com/xlab/vorbis-go/decoder"
	"github.com/xlab/vorbis-go/vorbis"

	"github.com/hajimehoshi/webmplayer/internal/libopus"
)

const samplesPerBuffer = 1024

type audioStream struct {
	codec             audioCodec
	channels          int
	samplingFrequency int

	src     <-chan webm.Packet
	packets []webm.Packet

	voDSP   vorbis.DspState
	voBlock vorbis.Block
	voPCM   [][][]float32

	opDecoder *libopus.Decoder
	opPCM     []float32

	frames []float32
}

type audioCodec string

const (
	audioCodecVorbis audioCodec = "A_VORBIS"
	audioCodecOpus   audioCodec = "A_OPUS"
)

func newAudioDecoder(codec audioCodec, codecPrivate []byte, channels, samplingFrequency int, src <-chan webm.Packet) (*audioStream, error) {
	d := &audioStream{
		channels:          channels,
		samplingFrequency: samplingFrequency,
		codec:             codec,
		src:               src,
	}
	switch codec {
	case audioCodecVorbis:
		var info vorbis.Info
		vorbis.InfoInit(&info)
		var comment vorbis.Comment
		vorbis.CommentInit(&comment)
		err := decoder.ReadHeaders(codecPrivate, &info, &comment)
		if err != nil {
			return nil, err
		}
		info.Deref()
		comment.Deref()
		if comment.Comments > 0 {
			comment.UserComments = make([][]byte, comment.Comments)
			comment.Deref()
		}
		if int(info.Channels) != channels {
			d.channels = int(channels)
			return nil, fmt.Errorf("webmplayer: channel count doesn't match: %d vs %d", info.Channels, channels)
		}
		if int(info.Rate) != samplingFrequency {
			d.samplingFrequency = int(info.Rate)
			return nil, fmt.Errorf("webmplayer: sample rate doesn't match: %d vs %d", info.Rate, samplingFrequency)
		}
		ret := vorbis.SynthesisInit(&d.voDSP, &info)
		if ret != 0 {
			return nil, fmt.Errorf("webmplayer: vorbis.SynthesisInit failed: %d", ret)
		}
		d.voPCM = [][][]float32{
			make([][]float32, channels),
		}
		vorbis.BlockInit(&d.voDSP, &d.voBlock)
		return d, nil
	case audioCodecOpus:
		var err error
		d.opDecoder, err = libopus.DecoderCreate(samplingFrequency, channels)
		if err != nil {
			return nil, err
		}
		d.opPCM = make([]float32, samplesPerBuffer*channels)
		return d, nil
	default:
		return d, fmt.Errorf("webmplayer: unsupported audio codec: %s", codec)
	}
}

func (a *audioStream) Read(buf []byte) (int, error) {
readFrames:
	if len(a.frames) > 0 {
		n := copy(unsafe.Slice((*float32)(unsafe.Pointer(unsafe.SliceData(buf))), len(buf)/4), a.frames)
		a.frames = a.frames[n:]
		return 4 * n, nil
	}

	for len(a.packets) == 0 {
		pkt, ok := <-a.src
		if !ok {
			n := min(len(buf)/4*4, 256)
			for i := range n {
				buf[i] = 0
			}
			return n, nil
		}
		if len(pkt.Data) == 0 {
			continue
		}
		a.packets = append(a.packets, pkt)
	}

	pkt := a.packets[0]
	a.packets = a.packets[1:]

	switch a.codec {
	case audioCodecVorbis:
		packet := &vorbis.OggPacket{
			Packet: pkt.Data,
			Bytes:  len(pkt.Data),
		}
		if ret := vorbis.Synthesis(&a.voBlock, packet); ret != 0 {
			return 0, fmt.Errorf("webmplayer: vorbis.Synthesis failed: %d", ret)
		}

		vorbis.SynthesisBlockin(&a.voDSP, &a.voBlock)

		sampleCount := vorbis.SynthesisPcmout(&a.voDSP, a.voPCM)
		if sampleCount == 0 {
			vorbis.SynthesisRead(&a.voDSP, sampleCount)
			return 0, nil
		}

		for ; sampleCount > 0; sampleCount = vorbis.SynthesisPcmout(&a.voDSP, a.voPCM) {
			for i := 0; i < int(sampleCount); i++ {
				for j := 0; j < a.channels; j++ {
					v := a.voPCM[0][j][:sampleCount][i]
					a.frames = append(a.frames, v)
					if a.channels == 1 {
						a.frames = append(a.frames, v)
					}
				}
			}
			vorbis.SynthesisRead(&a.voDSP, sampleCount)
		}

		goto readFrames

	case audioCodecOpus:
		sampleCount := a.opDecoder.DecodeFloat(pkt.Data, a.opPCM, 0)
		if sampleCount <= 0 {
			return 0, nil
		}

		origLen := len(a.frames)
		a.frames = append(a.frames, a.opPCM[:int(sampleCount)*a.channels]...)
		if a.channels == 1 {
			a.frames = append(a.frames, make([]float32, sampleCount)...)
			frames := a.frames[origLen:]
			for i := int(sampleCount) - 1; i > 0; i-- {
				frames[2*i] = frames[i]
				frames[2*i+1] = frames[i]
			}
		}

		goto readFrames

	default:
		return 0, fmt.Errorf("webmplayer: unsupported audio codec: %s", a.codec)
	}
}

func (a *audioStream) Channels() int {
	return a.channels
}

func (a *audioStream) SamplingFrequency() int {
	return a.samplingFrequency
}
