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
	op := &cgen.GenerateOptions{
		ProjectName: "libopus",
		TarGzURL:    "https://downloads.xiph.org/releases/opus/opus-1.5.2.tar.gz",
		TopDirs: []string{
			"include",
			"src",
		},
		AllowedFiles: []string{
			"COPYING",
			"README",
		},
		BlockedFiles: []string{
			"celt/opus_custom_demo.c",
			"src/opus_compare.c", // main function is defined.
			"src/opus_demo.c",
			"src/repacketizer_demo.c",
		},
		BlockedDirs: []string{
			"celt/dump_modes",
			"celt/tests",
			"cmake",
			"celt/arm",
			"celt/dump_modes",
			"celt/mips",
			"celt/tests",
			"celt/x86",
			"dnn",
			"doc",
			"silk/arm",
			"silk/fixed",
			"silk/float/x86",
			"silk/mips",
			"silk/tests",
			"silk/x86",
			"tests",
		},
	}
	if err := cgen.Generate(op); err != nil {
		return err
	}
	return nil
}
