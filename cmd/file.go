package main

import (
	"fmt"
	"os"
	"strings"
)

// FileHandler handles list file.
type FileHandler struct {
	file string
}

// NewFileHandler returns initialized *FileHandler
func NewFileHandler(file string) (*FileHandler, error) {
	info, err := os.Stat(file)
	if err == nil && info.IsDir() {
		return nil, fmt.Errorf("'%s' is dir, please set file path", file)
	}

	return &FileHandler{
		file: file,
	}, nil
}

// WriteAll writes lines into file
func (f *FileHandler) WriteAll(lines []string) error {
	fp, err := os.Create(f.file)
	if err != nil {
		return err
	}
	defer fp.Close()

	if _, err := fp.WriteString(strings.Join(lines, "\n")); err != nil {
		return err
	}
	return fp.Sync()
}
