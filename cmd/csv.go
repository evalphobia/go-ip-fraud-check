package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

// CSVHandler handles CSV file.
type CSVHandler struct {
	header    []string
	headerMap map[string]int
	reader    *csv.Reader
}

// NewCSVHandler returns initialized *CSVHandler
func NewCSVHandler(file string) (*CSVHandler, error) {
	info, err := os.Stat(file)
	if err == nil && info.IsDir() {
		return nil, fmt.Errorf("'%s' is dir, please set file path", file)
	}

	// load file
	fp, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	// load csv
	reader := csv.NewReader(fp)
	reader.LazyQuotes = true
	switch filepath.Ext(file) {
	case ".tsv":
		reader.Comma = '\t'
	}

	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	headerMap := make(map[string]int)
	for i, col := range header {
		headerMap[col] = i
	}

	return &CSVHandler{
		header:    header,
		headerMap: headerMap,
		reader:    reader,
	}, nil
}

// ReadAll reads all lines from CSV file.
func (f *CSVHandler) ReadAll() ([]map[string]string, error) {
	if f.reader == nil {
		return nil, fmt.Errorf("f.reader is nil")
	}

	lines, err := f.reader.ReadAll()
	if err != nil {
		return nil, err
	}

	header := f.header
	result := make([]map[string]string, len(lines))
	for i, line := range lines {
		r := make(map[string]string)
		for j, col := range line {
			r[header[j]] = col
		}
		result[i] = r
	}
	return result, nil
}
