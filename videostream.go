// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 Hajime Hoshi

package webmplayer

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ebml-go/webm"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xlab/libvpx-go/vpx"
)

type videoStream struct {
	src   <-chan webm.Packet
	ctx   *vpx.CodecCtx
	iface *vpx.CodecIface

	offscreen *ebiten.Image

	pos atomic.Int64

	err atomic.Pointer[error]

	m sync.Mutex
}

type videoCodec string

const (
	videoCodecVP8  videoCodec = "V_VP8"
	videoCodecVP9  videoCodec = "V_VP9"
	videoCodecVP10 videoCodec = "V_VP10"
)

func newVideoStream(codec videoCodec, src <-chan webm.Packet) (*videoStream, error) {
	v := &videoStream{
		src: src,
		ctx: vpx.NewCodecCtx(),
	}
	switch codec {
	case videoCodecVP8:
		v.iface = vpx.DecoderIfaceVP8()
	case videoCodecVP9:
		v.iface = vpx.DecoderIfaceVP9()
	default:
		return nil, fmt.Errorf("webmplayer: unsupported VPX codec: %s", codec)
	}
	if err := vpx.Error(vpx.CodecDecInitVer(v.ctx, v.iface, nil, 0, vpx.DecoderABIVersion)); err != nil {
		return nil, err
	}
	go v.loop()
	return v, nil
}

func (v *videoStream) Update(position time.Duration) error {
	if err := v.err.Load(); err != nil {
		return *err
	}
	v.pos.Store(int64(position))
	return nil
}

func (v *videoStream) Draw(f func(*ebiten.Image)) {
	v.m.Lock()
	defer v.m.Unlock()
	if v.offscreen == nil {
		return
	}
	f(v.offscreen)
}

func (v *videoStream) loop() {
loop:
	for pkt := range v.src {
		dataSize := uint32(len(pkt.Data))
		if err := vpx.Error(vpx.CodecDecode(v.ctx, string(pkt.Data), dataSize, nil, 0)); err != nil {
			v.err.Store(&err)
			return
		}
		pos := time.Duration(v.pos.Load())
		if pos-time.Second/60 > pkt.Timecode {
			continue loop
		}

		var iter vpx.CodecIter
		for img := vpx.CodecGetFrame(v.ctx, &iter); img != nil; img = vpx.CodecGetFrame(v.ctx, &iter) {
			img.Deref()
			if pos < pkt.Timecode {
				time.Sleep(pkt.Timecode - pos)
			}
			// TODO: Use img.ImageYCbCr and a shader.
			img := img.ImageRGBA()

			v.m.Lock()
			if v.offscreen != nil && v.offscreen.Bounds() != img.Bounds() {
				v.offscreen.Deallocate()
				v.offscreen = nil
			}
			if v.offscreen == nil {
				v.offscreen = ebiten.NewImage(img.Bounds().Dx(), img.Bounds().Dy())
			}
			v.offscreen.WritePixels(img.Pix)
			v.m.Unlock()
		}
	}
}
