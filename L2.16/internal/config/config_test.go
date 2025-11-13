package config_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"L2.16/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {

	def := config.Default()

	assert.True(t, def.Async, "expected Async true")
	assert.Equal(t, 5, def.Parallelism, "expected Parallelism 5")
	assert.Equal(t, 500*time.Millisecond, def.Delay, "expected Delay 500ms")
	assert.Equal(t, "RealPerson/1.0", def.UserAgent, "unexpected UserAgent")

	expected := []string{
		"Async", "Parallelism", "Delay", "RandomDelay", "UserAgent",
		"RequestTimeout", "IgnoreRobotsTxt",
	}

	rt := reflect.TypeOf(*def)
	for _, f := range expected {
		_, ok := rt.FieldByName(f)
		assert.True(t, ok, "missing field %s in Config", f)
	}

}

func TestLoad_NoSelectors_ReturnsErrNoSelectors(t *testing.T) {

	tmpDir := t.TempDir()
	origWD, _ := os.Getwd()
	defer func() { _ = os.Chdir(origWD) }()

	assert.NoError(t, os.Chdir(tmpDir), "failed to chdir")

	content := `
user_agent: "test-agent"
parallelism: 2
request_timeout: "1m"
download_html: false
download_css: false
download_scripts: false
download_images: false
download_videos: false
download_audio: false
download_iframes: false
download_fonts: false
download_icons: false
download_manifests: false
download_json: false
`
	assert.NoError(t, os.WriteFile(filepath.Join(tmpDir, "config.yaml"), []byte(content), 0777), "failed to write config")

	_, err := config.Load()
	assert.Error(t, err, "expected error for no selectors")
	assert.Equal(t, config.ErrNoSelectors, err, "expected ErrNoSelectors")

}

func TestLoad_ValidConfig(t *testing.T) {

	tmpDir := t.TempDir()
	origWD, _ := os.Getwd()
	defer func() { _ = os.Chdir(origWD) }()

	assert.NoError(t, os.Chdir(tmpDir), "failed to chdir")

	content := `
user_agent: "aboba"
parallelism: 3
request_timeout: "2m"
download_html: true
download_css: true
`
	assert.NoError(t, os.WriteFile(filepath.Join(tmpDir, "config.yaml"), []byte(content), 0777), "failed to write config")

	cfg, err := config.Load()
	assert.NoError(t, err, "Load returned error")
	assert.Equal(t, "aboba", cfg.UserAgent, "UserAgent mismatch")
	assert.Equal(t, 3, cfg.Parallelism, "Parallelism mismatch")
	assert.Equal(t, 2*time.Minute, cfg.RequestTimeout, "RequestTimeout mismatch")
	assert.True(t, cfg.DownloadHTML, "expected DownloadHTML enabled")
	assert.True(t, cfg.DownloadCSS, "expected DownloadCSS enabled")

}
