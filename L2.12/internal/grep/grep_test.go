package grep

import (
	"bytes"
	"os"
	"regexp"
	"testing"

	"L2.12/internal/buffer"
	"L2.12/internal/input"
)

func TestHighlightMatch(t *testing.T) {

	tests := []struct {
		input    string
		expected string
	}{
		{"qweqwe", bold + red + "qweqwe" + reset},
		{"123qwe", bold + red + "123qwe" + reset},
		{"", bold + red + "" + reset},
	}

	for _, tt := range tests {
		result := string(highlightMatch([]byte(tt.input)))
		if result != tt.expected {
			t.Errorf("highlightMatch(%q) = %q, expected = %q", tt.input, result, tt.expected)
		}
	}

}

func TestNeedsSeparator(t *testing.T) {

	tests := []struct {
		curr, last, A, B int
		expected         bool
	}{
		{5, 0, 1, 1, false},
		{5, 1, 0, 0, false},
		{10, 5, 1, 2, true},
		{5, 4, 1, 1, false},
	}

	for _, tt := range tests {

		data := &input.Data{
			CurrLineNumber:        tt.curr,
			LastPrintedLineNumber: tt.last,
			Flags: input.Flags{
				FlagA: tt.A,
				FlagB: tt.B,
			},
		}

		result := needsSeparator(data)
		if result != tt.expected {
			t.Errorf("(%d,%d,A=%d,B=%d) = %t, expected = %t", tt.curr, tt.last, tt.A, tt.B, result, tt.expected)
		}

	}

}

func captureOutput(f func()) string {

	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w

	f()

	if err := w.Close(); err != nil {
		os.Stdout = old
		return ""
	}
	os.Stdout = old

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		return ""
	}

	return buf.String()

}

func TestHandleInvertedMatch(t *testing.T) {

	data := &input.Data{
		Pattern:        regexp.MustCompile("cool pattern"),
		Flags:          input.Flags{Flagn: true},
		LinesBefore:    buffer.NewBuffer(1),
		CurrLineNumber: 1,
	}

	output := captureOutput(func() {
		handleInvertedMatch("cool line", 1, data)
	})

	if len(output) == 0 {
		t.Errorf("Expected output, got empty string")
	}

}

func TestHandleMatch(t *testing.T) {

	data := &input.Data{
		Pattern:        regexp.MustCompile("pAtTeRn"),
		Flags:          input.Flags{Flagn: true},
		CurrLineNumber: 1,
		LinesBefore:    buffer.NewBuffer(1),
	}

	line := "pa-ta-tern"

	output := captureOutput(func() {
		handleMatch(line, data)
	})

	if !data.MatchFound {
		t.Errorf("Expected MatchFound=true, got false")
	}

	if len(output) == 0 {
		t.Errorf("Expected printed output, got empty string")
	}

}

func TestPrintLine(t *testing.T) {

	tests := []struct {
		name       string
		match      bool
		multiFile  bool
		flags      input.Flags
		expectedIn string
	}{
		{"single file no line numbers", true, false, input.Flags{}, "abobus"},
		{"multi file with line numbers", false, true, input.Flags{Flagn: true}, "file.txt"},
	}

	for _, tt := range tests {

		data := &input.Data{
			Flags:       tt.flags,
			FileNames:   []string{},
			CurrentFile: "file.txt",
		}

		if tt.multiFile {
			data.FileNames = []string{"file1", "file2"}
		}

		output := captureOutput(func() {
			printLine(tt.match, data, "abobus", 1, ":")
		})

		if len(output) == 0 || output == "\n" {
			t.Errorf("%s: expected printed line, got empty", tt.name)
		}

		if tt.expectedIn != "" && !bytes.Contains([]byte(output), []byte(tt.expectedIn)) {
			t.Errorf("%s: output %q does not contain %q", tt.name, output, tt.expectedIn)
		}

	}
}
