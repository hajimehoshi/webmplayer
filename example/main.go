// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 Hajime Hoshi

package main

import (
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/webmplayer"
)

func main() {
	flag.Parse()
	if err := xmain(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func xmain() error {
	streams := make([]io.ReadSeeker, 0, 2)
	for _, opt := range flag.Args() {
		f, err := os.Open(opt)
		if err != nil {
			return err
		}
		streams = append(streams, f)
		if len(streams) >= 2 {
			break
		}
	}

	player, err := webmplayer.NewPlayer(streams...)
	if err != nil {
		return err
	}

	if player.VideoCodecID() != "" {
		w, h := player.VideoSize()
		slog.Info("Video",
			"codec", player.VideoCodecID(),
			"width", w,
			"height", h,
			"duration", player.VideoDuration())
	}
	if player.AudioCodecID() != "" {
		slog.Info("Audio",
			"codec", player.AudioCodecID(),
			"channels", player.AudioChannels(),
			"samplingFrequency", player.AudioSamplingFrequency(),
			"duration", player.AudioDuration())
	}

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("WebM Player")
	game := NewGame(player)
	if err := ebiten.RunGame(game); err != nil {
		return err
	}

	return nil
}

type Game struct {
	player *webmplayer.Player
}

func NewGame(p *webmplayer.Player) *Game {
	return &Game{
		player: p,
	}
}

func (g *Game) Update() error {
	return g.player.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	w, h := g.player.VideoSize()
	if w == 0 || h == 0 {
		return
	}

	op := &webmplayer.PlayerDrawOptions{}
	scale := min(float64(screen.Bounds().Dx())/float64(w), float64(screen.Bounds().Dy())/float64(h))
	op.GeoM.Scale(scale, scale)
	g.player.Draw(screen, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
