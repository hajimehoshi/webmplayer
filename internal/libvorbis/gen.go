// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 Hajime Hoshi

//go:build ignore

package main

import (
	"log/slog"
	"os"

	"github.com/hajimehoshi/webmplayer/internal/cgen"
)

func main() {
	if err := xmain(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func xmain() error {
	oggOp := &cgen.GenerateOptions{
		ProjectName: "libogg",
		TarGzURL:    "https://downloads.xiph.org/releases/ogg/libogg-1.3.5.tar.gz",
		TopDirs: []string{
			"include",
			"src",
		},
		AllowedFiles: []string{
			"README.md",
		},
		BlockedFiles: []string{},
		BlockedDirs:  []string{},
	}

	vorbisOp := &cgen.GenerateOptions{
		ProjectName: "libvorbis",
		TarGzURL:    "https://downloads.xiph.org/releases/vorbis/libvorbis-1.3.7.tar.gz",
		TopDirs: []string{
			"include",
			"lib",
		},
		AllowedFiles: []string{
			"COPYING",
		},
		BlockedFiles: []string{
			"lib/psytune.c", // a dead code.
			"lib/barkmel.c",
			"lib/tone.c",
		},
		BlockedDirs: []string{
			"examples",
			"lib/modes",
			"symbian",
			"test",
			"vq",
		},
	}

	if err := cgen.Generate(oggOp, vorbisOp); err != nil {
		return err
	}
	return nil
}
