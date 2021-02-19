package main

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"path/filepath"
	"testing"
)

func TestAddWatermarks(t *testing.T, msg, inFile, outFile string, selectedPages []string, mode, modeParam, desc string, onTop bool) {
	t.Helper()
	inFile = filepath.Join(inDir, inFile)
	s := "watermark"
	if onTop {
		s = "stamp"
	}
	outFile = filepath.Join("./out", s, mode, outFile)

	var err error
	switch mode {
	case "text":
		err = api.AddTextWatermarksFile(inFile, outFile, selectedPages, onTop, modeParam, desc, nil)
	case "image":
		err = api.AddImageWatermarksFile(inFile, outFile, selectedPages, onTop, modeParam, desc, nil)
	case "pdf":
		err = api.AddPDFWatermarksFile(inFile, outFile, selectedPages, onTop, modeParam, desc, nil)
	}
	if err != nil {
		t.Fatalf("%s %s: %v\n", msg, outFile, err)
	}
	if err := api.ValidateFile(outFile, nil); err != nil {
		t.Fatalf("%s: %v\n", msg, err)
	}
}

