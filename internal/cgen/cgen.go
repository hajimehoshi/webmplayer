// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024 Hajime Hoshi

package cgen

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

type GenerateOptions struct {
	ProjectName  string
	TarGzURL     string
	TopDirs      []string
	AllowedFiles []string
	BlockedFiles []string
	BlockedDirs  []string
}

type context struct {
	options *GenerateOptions
}

type entry struct {
	name    string
	content []byte
	context *context
}

func Generate(options ...*GenerateOptions) error {
	suffixes := make([]string, 0, len(options))
	for _, op := range options {
		suffixes = append(suffixes, op.ProjectName)
	}
	if err := clean(suffixes); err != nil {
		return err
	}

	var entries []entry

	for _, op := range options {
		c := &context{
			options: op,
		}

		if err := c.fetchTarGz(); err != nil {
			return err
		}

		tarGzFileName, err := c.tarGzFileName()
		if err != nil {
			return err
		}
		f, err := os.Open(tarGzFileName)
		if err != nil {
			return err
		}
		defer f.Close()

		entries, err = c.appendEntriesFromTarGz(entries, f)
		if err != nil {
			return err
		}
	}

	if err := outputFiles(".", entries); err != nil {
		return err
	}

	return nil
}

func clean(suffixes []string) error {
	if err := filepath.Walk(".", func(p string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && p != "." {
			return filepath.SkipDir
		}

		remove := strings.HasSuffix(p, ".c") || strings.HasSuffix(p, ".h")
		if !remove {
			for _, suffix := range suffixes {
				if strings.HasSuffix(strings.TrimSuffix(p, path.Ext(p)), suffix) {
					remove = true
					break
				}
			}
		}

		if remove {
			return os.Remove(p)
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (c *context) fileNameSuffix() string {
	return "-" + c.options.ProjectName
}

func (c *context) tarGzFileName() (string, error) {
	u, err := url.Parse(c.options.TarGzURL)
	if err != nil {
		return "", err
	}
	return path.Base(u.Path), nil
}

func (c *context) fetchTarGz() error {
	tarGzFileName, err := c.tarGzFileName()
	if err != nil {
		return err
	}

	if _, err := os.Stat(tarGzFileName); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	} else if err == nil {
		return nil
	}

	res, err := http.Get(c.options.TarGzURL)
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

func (c *context) appendEntriesFromTarGz(entries []entry, src io.Reader) ([]entry, error) {
	s, err := gzip.NewReader(src)
	if err != nil {
		return nil, err
	}

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
			entries = append(entries, entry{
				name:    strings.Join(strings.Split(name, "/")[1:], "/"),
				content: bs,
				context: c,
			})
		default:
			return nil, fmt.Errorf("unsupported type: %v", header.Typeflag)
		}
	}

	return entries, nil
}

func (c *context) isAllowed(name string) bool {
	if strings.HasSuffix(name, ".c") {
		return true
	}
	if strings.HasSuffix(name, ".h") {
		return true
	}
	for _, name1 := range c.options.AllowedFiles {
		if name == name1 {
			return true
		}
	}
	return false
}

func outputFiles(dst string, entries []entry) error {
entries:
	for _, entry := range entries {
		if !entry.context.isAllowed(entry.name) {
			continue
		}

		for _, dir := range entry.context.options.BlockedDirs {
			if strings.HasPrefix(entry.name, dir+"/") {
				continue entries
			}
		}

		for _, name1 := range entry.context.options.BlockedFiles {
			if entry.name == name1 {
				continue entries
			}
		}

		bs := entry.content

		// Rewrite include paths.
		if strings.HasSuffix(entry.name, ".c") || strings.HasSuffix(entry.name, ".h") {
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

				var needReplace bool
				for _, entry1 := range entries {
					key := entry1.name
					for _, dir := range entry.context.options.TopDirs {
						key = strings.TrimPrefix(key, dir+"/")
					}
					if key == p {
						needReplace = true
						break
					}
					// Relative path.
					if strings.HasSuffix(key, "/"+p) {
						p = key
						needReplace = true
						break
					}
				}
				if needReplace {
					p = strings.ReplaceAll(p, "/", "_")
					newBS = append(newBS, []byte(m[1]+p+m[3])...)
				} else {
					newBS = append(newBS, line...)
				}
				newBS = append(newBS, '\n')
				continue
			}
			if err := s.Err(); err != nil {
				return err
			}
			bs = newBS
		}

		outName := entry.name
		if tokens := strings.Split(entry.name, "/"); len(tokens) > 1 {
			if slices.Contains(entry.context.options.TopDirs, tokens[0]) {
				tokens = tokens[1:]
			}
			outName = strings.Join(tokens, "_")
		}

		if slices.Contains(entry.context.options.AllowedFiles, outName) {
			ext := path.Ext(outName)
			outName = strings.TrimSuffix(outName, ext) + entry.context.fileNameSuffix() + ext
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
