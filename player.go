// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 Hajime Hoshi

package webmplayer

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type Player struct {
	width  int
	height int

	videoStream *videoStream
	audioStream *audioStream
	audioPlayer *audio.Player
}

func NewPlayer(streams ...io.ReadSeeker) (*Player, error) {
	stream1, stream2, err := discoverStreams(streams...)
	if err != nil {
		return nil, err
	}
	if stream1 == nil {
		return nil, fmt.Errorf("webmplayer: nothing to play")
	}

	videoStream := stream1.VideoStream()
	audioStream := stream1.AudioStream()
	if stream2 != nil {
		audioStream = stream2.AudioStream()
	}

	var w, h int
	if vtrack := stream1.Meta().FindFirstVideoTrack(); vtrack != nil {
		w, h = int(vtrack.DisplayWidth), int(vtrack.DisplayHeight)
	}

	v := &Player{
		width:       w,
		height:      h,
		videoStream: videoStream,
		audioStream: audioStream,
	}
	if audioStream != nil {
		ctx := audio.NewContext(audioStream.SampleRate())
		p, err := ctx.NewPlayerF32(audioStream)
		if err != nil {
			return nil, err
		}
		p.Play()
		v.audioPlayer = p
	}
	return v, nil
}

func (v *Player) VideoSize() (int, int) {
	return v.width, v.height
}

func (v *Player) Update() error {
	v.videoStream.Update(v.audioPlayer.Position())
	return nil
}

type PlayerDrawOptions struct {
	GeoM       ebiten.GeoM
	ColorScale ebiten.ColorScale
	Blend      ebiten.Blend
}

func (v *Player) Draw(screen *ebiten.Image, options *PlayerDrawOptions) {
	if v.videoStream == nil {
		return
	}
	v.videoStream.Draw(func(image *ebiten.Image) {
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

// discoverStreams returns both Video and Audio streams if in separate inputs,
// otherwise only the first stream would be returned (Video / Audio / Video + Audio).
func discoverStreams(streams ...io.ReadSeeker) (*stream, *stream, error) {
	if len(streams) == 0 {
		return nil, nil, fmt.Errorf("webmplayer: no streams found")
	}

	if len(streams) == 1 {
		stream, err := newStream(streams[0])
		if err != nil {
			return nil, nil, err
		}
		return stream, nil, nil
	}

	var stream1Video bool
	var stream1Audio bool
	stream1, err := newStream(streams[0])
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
	stream2, err := newStream(streams[1])
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
