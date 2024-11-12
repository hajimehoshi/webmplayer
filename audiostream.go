// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 Hajime Hoshi

package webmplayer

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/ebml-go/webm"
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
		info, comment, err := readVorbisCodecPrivate(codecPrivate)
		if err != nil {
			return nil, err
		}
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
		ret := vorbis.SynthesisInit(&d.voDSP, info)
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

func readVorbisCodecPrivate(codecPrivate []byte) (*vorbis.Info, *vorbis.Comment, error) {
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

	var info vorbis.Info
	vorbis.InfoInit(&info)
	var comment vorbis.Comment
	vorbis.CommentInit(&comment)

	for i := 0; i < 3; i++ {
		packet := vorbis.OggPacket{
			Bytes:  len(headers[i]),
			Packet: headers[i],
		}
		if i == 0 {
			packet.BOS = 1
		}
		if ret := vorbis.SynthesisHeaderin(&info, &comment, &packet); ret < 0 {
			return nil, nil, fmt.Errorf("webmplayer: %d. header damaged", i+1)
		}
	}

	info.Deref()
	comment.Deref()

	return &info, &comment, nil
}
