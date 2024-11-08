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

	stream1, stream2, err := discoverStreams(streams...)
	if err != nil {
		return err
	}
	if stream1 == nil {
		return fmt.Errorf("webmplayer: nothing to play")
	}

	videoStream := stream1.VideoStream()
	audioStream := stream1.AudioStream()
	if stream2 != nil {
		audioStream = stream2.AudioStream()
	}

	var player *Player
	var w, h int
	if vtrack := stream1.Meta().FindFirstVideoTrack(); vtrack != nil {
		w, h = int(vtrack.DisplayWidth), int(vtrack.DisplayHeight)
		p, err := NewPlayer(videoStream, audioStream)
		if err != nil {
			return err
		}
		player = p
	} else {
		p, err := NewPlayer(nil, audioStream)
		if err != nil {
			return err
		}
		player = p
	}

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("WebM Player")
	game := NewGame(player, w, h)
	if err := ebiten.RunGame(game); err != nil {
		return err
	}

	return nil
}

// discoverStreams returns both Video and Audio streams if in separate inputs,
// otherwise only the first stream would be returned (V/A/V+A).
func discoverStreams(streams ...io.ReadSeeker) (*Stream, *Stream, error) {
	if len(streams) == 0 {
		return nil, nil, fmt.Errorf("webmplayer: no streams found")
	}

	if len(streams) == 1 {
		stream, err := NewStream(streams[0])
		if err != nil {
			return nil, nil, err
		}
		return stream, nil, nil
	}

	var stream1Video bool
	var stream1Audio bool
	stream1, err := NewStream(streams[0])
	if err != nil {
		slog.Warn(err.Error())
	} else {
		stream1Video = stream1.Meta().FindFirstVideoTrack() != nil
		stream1Audio = stream1.Meta().FindFirstAudioTrack() != nil
	}
	if stream1Video && stream1Audio {
		// Found both Video+Audio in the first stream.
		return stream1, nil, nil
	}

	var stream2Video bool
	var stream2Audio bool
	stream2, err := NewStream(streams[1])
	if err != nil {
		slog.Warn(err.Error())
	} else {
		stream2Video = stream2.Meta().FindFirstVideoTrack() != nil
		stream2Audio = stream2.Meta().FindFirstAudioTrack() != nil
	}

	switch {
	case stream1Video && stream2Audio:
		// Took Video from the first stream, Audio from the second.
		return stream1, stream2, nil
	case stream1Audio && stream2Video:
		// Took Audio from the first stream, Video from the second.
		return stream2, stream1, nil
	case stream1Video:
		// Took Video from the first stream, no Audio found.
		return stream1, nil, nil
	case stream2Video:
		// Took Video from the second stream, no Audio found.
		return stream2, nil, nil
	case stream1Audio:
		// Took Audio from the first stream, no Video found.
		return stream1, nil, nil
	case stream2Audio:
		// Took Audio from the second stream, no Video found.
		return stream2, nil, nil
	default:
		// No Video or Audio found.
		return nil, nil, nil
	}
}

type Game struct {
	player *Player
	width  int
	height int
}

func NewGame(p *Player, width, height int) *Game {
	return &Game{
		player: p,
		width:  width,
		height: height,
	}
}

func (g *Game) Update() error {
	return g.player.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &PlayerDrawOptions{}
	scale := min(float64(screen.Bounds().Dx())/float64(g.width), float64(screen.Bounds().Dy())/float64(g.height))
	op.GeoM.Scale(scale, scale)
	g.player.Draw(screen, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
