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
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

const tarGzURL = "https://downloads.xiph.org/releases/opus/opus-1.5.2.tar.gz"

var tarGzFileName string

func init() {
	u, err := url.Parse(tarGzURL)
	if err != nil {
		panic(err)
	}
	tarGzFileName = path.Base(u.Path)
}

const projectName = "libopus"

const fileNameSuffix = "-" + projectName

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

	slog.Info("Fetching " + projectName)
	if err := fetchTarGz(); err != nil {
		return err
	}

	f, err := os.Open(tarGzFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	slog.Info("Extracting " + projectName)
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
		if !strings.HasSuffix(base, fileNameSuffix) &&
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

func fetchTarGz() error {
	_, err := os.Stat(tarGzFileName)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if err == nil {
		return nil
	}

	res, err := http.Get(tarGzURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	f, err := os.Create(tarGzFileName)
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
			"silk/float/x86",
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
			outName = outName + fileNameSuffix
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
