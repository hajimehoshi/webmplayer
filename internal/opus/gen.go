// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 Hajime Hoshi

//go:build ignore

package main

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

func main() {
	if err := xmain(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func xmain() error {
	slog.Info("Cleaning *.c and *.h files")
	if err := clean(); err != nil {
		return err
	}

	slog.Info("Fetching opus")
	if err := fetchOpus(); err != nil {
		return err
	}

	f, err := os.Open("opus-1.5.2.tar.gz")
	if err != nil {
		return err
	}
	defer f.Close()

	slog.Info("Extracting opus")
	entries, err := extractTarGz(f)
	if err != nil {
		return err
	}

	slog.Info("Outputting files")
	if err := outputFiles(".", entries); err != nil {
		return err
	}

	return nil
}

func clean() error {
	if err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != "." {
			return filepath.SkipDir
		}
		base := filepath.Base(path)
		if !strings.HasSuffix(base, "-opus") &&
			!strings.HasSuffix(base, ".c") &&
			!strings.HasSuffix(base, ".h") {
			return nil
		}
		return os.Remove(base)
	}); err != nil {
		return err
	}
	return nil
}

func fetchOpus() error {
	_, err := os.Stat("opus-1.5.2.tar.gz")
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if err == nil {
		return nil
	}

	res, err := http.Get("https://downloads.xiph.org/releases/opus/opus-1.5.2.tar.gz")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	f, err := os.Create("opus-1.5.2.tar.gz")
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	if _, err := io.Copy(w, res.Body); err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}
	return nil
}

func extractTarGz(src io.Reader) (map[string][]byte, error) {
	s, err := gzip.NewReader(src)
	if err != nil {
		return nil, err
	}

	entries := map[string][]byte{}
	r := tar.NewReader(s)
	for {
		header, err := r.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		switch header.Typeflag {
		case tar.TypeDir:
			continue
		case tar.TypeReg:
			name := header.Name
			bs, err := io.ReadAll(r)
			if err != nil {
				return nil, err
			}
			entries[strings.Join(strings.Split(name, "/")[1:], "/")] = bs
		default:
			return nil, fmt.Errorf("unsupported type: %v", header.Typeflag)
		}
	}

	return entries, nil
}

func outputFiles(dst string, entries map[string][]byte) error {
entries:
	for name, bs := range entries {
		if name != "README" &&
			name != "COPYING" &&
			!strings.HasSuffix(name, ".c") &&
			!strings.HasSuffix(name, ".h") {
			continue
		}
		if name == "LICENSE_PLEASE_READ.txt" {
			println(name)
		}

		for _, name1 := range []string{
			"dnn/dump_data.c",            // main function is defined.
			"dnn/write_lpcnet_weights.c", // main function is defined.
			"src/opus_compare.c",         // main function is defined.
		} {
			if name == name1 {
				continue entries
			}
		}
		for _, dir := range []string{
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
			"silk/mips",
			"silk/tests",
			"silk/x86",
			"tests"} {
			if strings.HasPrefix(name, dir+"/") {
				continue entries
			}
		}
		if strings.HasSuffix(name, "_demo.c") {
			continue
		}

		// Add go:build directives if needed.
		if strings.HasSuffix(name, ".c") || strings.HasSuffix(name, ".h") {
			// included by dnn/lossgen.c.
			if name == "dnn/parse_lpcnet_weights.c" {
				bs = append([]byte("//go:build ignore\n\n"), bs...)
			} else {
				pathTokens := strings.Split(strings.TrimSuffix(name, path.Ext(name)), "/")
				var tokens []string
				for _, pathToken := range pathTokens {
					tokens = append(tokens, strings.Split(pathToken, "_")...)
				}
				switch {
				case slices.Contains(tokens, "FIX") || slices.Contains(tokens, "dotprod") || slices.Contains(tokens, "ne10") || slices.Contains(tokens, "neon") || slices.Contains(tokens, "osce"):
					bs = append([]byte("//go:build ignore\n\n"), bs...)
				case slices.Contains(tokens, "arm") || slices.Contains(tokens, "neon"):
					bs = append([]byte("//go:build arm || arm64\n\n"), bs...)
				case slices.Contains(tokens, "armv4"), slices.Contains(tokens, "armv5e"):
					bs = append([]byte("//go:build arm\n\n"), bs...)
				case slices.Contains(tokens, "arm64"):
					bs = append([]byte("//go:build arm64\n\n"), bs...)
				case slices.Contains(tokens, "mips") || slices.Contains(tokens, "mipsr1"):
					bs = append([]byte("//go:build mips || mips64\n\n"), bs...)
				case slices.Contains(tokens, "avx") || slices.Contains(tokens, "avx2") ||
					slices.Contains(tokens, "sse") || slices.Contains(tokens, "sse2") || slices.Contains(tokens, "sse4") ||
					slices.Contains(tokens, "x86") || slices.Contains(tokens, "x86cpu"):
					bs = append([]byte("//go:build 386 || amd64\n\n"), bs...)
				}
			}
		}

		// Rewrite include paths.
		if strings.HasSuffix(name, ".c") || strings.HasSuffix(name, ".h") {
			reInclude := regexp.MustCompile(`^(\s*#\s*include\s+["<])(.*)([">])$`)
			var newBS []byte
			s := bufio.NewScanner(bytes.NewReader(bs))
			for s.Scan() {
				line := s.Text()
				m := reInclude.FindStringSubmatch(line)
				if m == nil {
					newBS = append(newBS, line...)
					newBS = append(newBS, '\n')
					continue
				}

				p := m[2]
				for strings.HasPrefix(p, "../") {
					p = strings.TrimPrefix(p, "../")
				}
				for key := range entries {
					key = strings.TrimPrefix(key, "src/")
					key = strings.TrimPrefix(key, "include/")
					if key == p {
						break
					}
					if strings.HasSuffix(key, "/"+p) {
						p = key
						break
					}
				}
				p = strings.ReplaceAll(p, "/", "_")
				newBS = append(newBS, []byte(m[1]+p+m[3])...)
				newBS = append(newBS, '\n')
				continue
			}
			if err := s.Err(); err != nil {
				return err
			}
			bs = newBS
		}

		outName := name
		if path.Dir(name) != "." {
			tokens := strings.Split(name, "/")
			if tokens[0] == "src" || tokens[0] == "include" {
				tokens = tokens[1:]
			}
			outName = strings.Join(tokens, "_")
		}

		if filepath.Ext(outName) == "" {
			outName = outName + "-opus"
		}
		if _, err := os.Stat(filepath.Join(dst, outName)); err == nil {
			return fmt.Errorf("file already exists: %s", outName)
		}
		if err := os.WriteFile(filepath.Join(dst, outName), bs, 0644); err != nil {
			return err
		}
	}
	return nil
}
