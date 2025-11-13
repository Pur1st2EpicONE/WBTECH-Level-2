package utils_test

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"

	"L2.16/internal/config"
	"L2.16/internal/utils"

	"github.com/stretchr/testify/assert"
)

func TestToHumanSize(t *testing.T) {

	cases := []struct {
		in  int64
		out string
	}{
		{512, "512B"},
		{2048, "2K"},
		{5 * 1024 * 1024, "5.0M"},
	}

	for _, c := range cases {
		got := utils.ToHumanSize(c.in)
		assert.Equal(t, c.out, got, "ToHumanSize(%d) mismatch", c.in)
	}

}

func TestBuildSelector(t *testing.T) {

	cfg := &config.Config{
		DownloadHTML:   true,
		DownloadCSS:    true,
		DownloadImages: true,
		DownloadJSON:   true,
	}

	selector := utils.BuildSelector(cfg)
	assert.Contains(t, selector, "a[href]", "selector missing a[href]")
	assert.Contains(t, selector, "link[rel=stylesheet][href]", "selector missing stylesheet link")
	assert.Contains(t, selector, "img[src]", "selector missing img[src]")
	assert.Contains(t, selector, "a[href$='.json']", "selector missing json link")

}

func TestWgetHeaderAndFormatOutput(t *testing.T) {

	old := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	utils.WgetHeader("http://aboba.com/test", "/tmp/example", "text/html", 1234)
	utils.FormatOutput(3, 2048, 2*time.Second)

	_ = writer.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, reader)
	out := buf.String()

	assert.Contains(t, out, "http://aboba.com/test", "expected URL in output")
	assert.Contains(t, out, "Saving to:", "expected 'Saving to' in output")
	assert.Contains(t, out, "Downloaded: 3 files", "expected downloads summary in output")

}
