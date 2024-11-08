// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 Hajime Hoshi

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type Player struct {
	videoDecoder *VideoStream
	audioDecoder *AudioStream
	audioPlayer  *audio.Player
}

func NewPlayer(videoDecoder *VideoStream, audioDecoder *AudioStream) (*Player, error) {
	v := &Player{
		videoDecoder: videoDecoder,
		audioDecoder: audioDecoder,
	}
	if audioDecoder != nil {
		ctx := audio.NewContext(audioDecoder.SampleRate())
		p, err := ctx.NewPlayerF32(audioDecoder)
		if err != nil {
			return nil, err
		}
		p.Play()
		v.audioPlayer = p
	}
	return v, nil
}

func (v *Player) Update() error {
	v.videoDecoder.Update(v.audioPlayer.Position())
	return nil
}

type PlayerDrawOptions struct {
	GeoM       ebiten.GeoM
	ColorScale ebiten.ColorScale
	Blend      ebiten.Blend
}

func (v *Player) Draw(screen *ebiten.Image, options *PlayerDrawOptions) {
	if v.videoDecoder == nil {
		return
	}
	v.videoDecoder.Draw(func(image *ebiten.Image) {
		op := &ebiten.DrawImageOptions{}
		op.Filter = ebiten.FilterLinear
		if options != nil {
			op.GeoM = options.GeoM
			op.ColorScale = options.ColorScale
			op.Blend = options.Blend
		}
		screen.DrawImage(image, op)
	})
}
