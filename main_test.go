package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	var mockStdOut bytes.Buffer
	if err := run(inputFile, "", &mockStdOut, true); err != nil {
		t.Fatal(err)
	}
	resultFile := strings.TrimSpace(mockStdOut.String())

	result, err := os.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	result = bytes.ReplaceAll(result, []byte("\r\n"), []byte("\n"))
	expected = bytes.ReplaceAll(expected, []byte("\r\n"), []byte("\n"))

	if !bytes.Equal(expected, result) {
		t.Logf("golden: %q", expected)
		t.Logf("result: %q", result)
		t.Error("Result content does not match golden file")
	}
}

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}
	result, err := parseContent(input, "")
	if err != nil {
		t.Fatal(err)
	}
	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	result = bytes.ReplaceAll(result, []byte("\r\n"), []byte("\n"))
	expected = bytes.ReplaceAll(expected, []byte("\r\n"), []byte("\n"))

	if !bytes.Equal(expected, result) {
		t.Logf("golden:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Error("Result content does not match golden file")
	}
}
