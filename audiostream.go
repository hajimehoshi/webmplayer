// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 Hajime Hoshi

package webmplayer

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/ebml-go/webm"

	"github.com/hajimehoshi/webmplayer/internal/libopus"
	"github.com/hajimehoshi/webmplayer/internal/libvorbis"
)

const samplesPerBuffer = 1024

type audioStream struct {
	codec             audioCodec
	channels          int
	samplingFrequency int

	src     <-chan webm.Packet
	packets []webm.Packet

	// voInfo must be kept as voDPS has a reference to it.
	voInfo  *libvorbis.Info
	voDSP   *libvorbis.DspState
	voBlock *libvorbis.Block

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
	a := &audioStream{
		channels:          channels,
		samplingFrequency: samplingFrequency,
		codec:             codec,
		src:               src,
	}
	// TODO: Clear vo* and op* objects explicitly when a is finalized.
	switch codec {
	case audioCodecVorbis:
		info, _, err := readVorbisCodecPrivate(codecPrivate)
		if err != nil {
			return nil, err
		}
		a.voInfo = info

		if info.Channels() != channels {
			a.channels = int(channels)
			return nil, fmt.Errorf("webmplayer: channel count doesn't match: %d vs %d", info.Channels(), channels)
		}
		if info.Rate() != samplingFrequency {
			a.samplingFrequency = info.Rate()
			return nil, fmt.Errorf("webmplayer: sample rate doesn't match: %d vs %d", info.Rate(), samplingFrequency)
		}

		dsp, err := libvorbis.SynthesisInit(info)
		if err != nil {
			return nil, fmt.Errorf("webmplayer: libvorbis.SynthesisInit failed: %w", err)
		}
		a.voDSP = dsp

		block, err := libvorbis.BlockInit(a.voDSP)
		if err != nil {
			return nil, fmt.Errorf("webmplayer: libvorbis.BlockInit failed: %w", err)
		}
		a.voBlock = block

		return a, nil

	case audioCodecOpus:
		var err error
		a.opDecoder, err = libopus.DecoderCreate(samplingFrequency, channels)
		if err != nil {
			return nil, err
		}
		a.opPCM = make([]float32, samplesPerBuffer*channels)
		return a, nil
	default:
		return a, fmt.Errorf("webmplayer: unsupported audio codec: %s", codec)
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
		packet := &libvorbis.OggPacket{
			Packet: pkt.Data,
		}
		if err := libvorbis.Synthesis(a.voBlock, packet); err != nil {
			return 0, fmt.Errorf("webmplayer: libvorbis.Synthesis failed: %w", err)
		}

		if err := libvorbis.SynthesisBlockin(a.voDSP, a.voBlock); err != nil {
			return 0, fmt.Errorf("webmplayer: libvorbis.SynthesisBlockin failed: %w", err)
		}

		for pcm := libvorbis.SynthesisPcmout(a.voDSP); len(pcm) > 0 && len(pcm[0]) > 0; pcm = libvorbis.SynthesisPcmout(a.voDSP) {
			switch a.channels {
			case 1:
				for i := range pcm[0] {
					v := pcm[0][i]
					a.frames = append(a.frames, v, v)
				}
			case 2:
				for i := range pcm[0] {
					for ch := range pcm {
						v := pcm[ch][i]
						a.frames = append(a.frames, v)
					}
				}
			default:
				return 0, fmt.Errorf("webmplayer: unsupported channel count: %d", a.channels)
			}
			if err := libvorbis.SynthesisRead(a.voDSP, len(pcm[0])); err != nil {
				return 0, fmt.Errorf("webmplayer: libvorbis.SynthesisRead failed: %w", err)
			}
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

func readVorbisCodecPrivate(codecPrivate []byte) (*libvorbis.Info, *libvorbis.Comment, error) {
	if len(codecPrivate) < 1 {
		return nil, nil, errors.New("webmplayer: codec private data is too short")
	}

	p := codecPrivate

	// https://www.matroska.org/technical/codec_specs.html
	// > Byte 1: number of distinct packets #p minus one inside the CodecPrivate block. This MUST be “2” for current (as of 2016-07-08) Vorbis headers.
	if p[0] != 0x02 {
		return nil, nil, fmt.Errorf("webmplayer: wrong codec private data for Vorbis: %d", p[0])
	}
	offset := 1
	p = p[1:]

	headers := make([][]byte, 3)
	var size0, size1 int

	// https://xiph.org/vorbis/doc/framing.html
	// > The raw packet is logically divided into [n] 255 byte segments and a last fractional segment of < 255 bytes.
	// > A packet size may well consist only of the trailing fractional segment, and a fractional segment may be zero length.
	// > These values, called "lacing values" are then saved and placed into the header segment table.
	for i := 0; i < 2; i++ {
		for (p[0] == 0xff) && offset < len(codecPrivate) {
			if i == 0 {
				size0 += 0xff
			} else {
				size1 += 0xff
			}
			offset++
			p = p[1:]
		}
		if offset >= len(codecPrivate)-1 {
			return nil, nil, errors.New("webmplayer: header sizes damaged")
		}
		if i == 0 {
			size0 += int(p[0])
		} else {
			size1 += int(p[0])
		}
		offset++
		p = p[1:]
	}
	headers[0] = codecPrivate[offset : offset+size0]
	headers[1] = codecPrivate[offset+size0 : offset+size0+size1]
	headers[2] = codecPrivate[offset+size0+size1:]

	info := libvorbis.InfoInit()
	comment := libvorbis.CommentInit()

	for i := 0; i < 3; i++ {
		packet := &libvorbis.OggPacket{
			Packet: headers[i],
			BOS:    i == 0,
		}
		if err := libvorbis.SynthesisHeaderin(info, comment, packet); err != nil {
			return nil, nil, fmt.Errorf("webmplayer: libvorbis.SynthesisHeaderin failed: %w", err)
		}
	}

	return info, comment, nil
}
