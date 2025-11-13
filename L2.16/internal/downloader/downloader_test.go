package downloader_test

import (
	"os"
	"path/filepath"
	"testing"

	"L2.16/internal/downloader"
	"github.com/stretchr/testify/assert"
)

func TestSave_CreatesDirAndFile(t *testing.T) {

	tmp := t.TempDir()
	target := filepath.Join(tmp, "subdir", "file.txt")
	data := []byte("hello world")

	res := downloader.Save(target, data)
	assert.NoError(t, res.Err, "Save returned error")
	assert.Equal(t, len(data), res.Size, "Size mismatch")

	info, err := os.Stat(target)
	assert.NoError(t, err, "stat target failed")
	assert.Equal(t, int64(len(data)), info.Size(), "file size mismatch")

	content, err := os.ReadFile(target)
	assert.NoError(t, err, "read file failed")
	assert.Equal(t, string(data), string(content), "file content mismatch")

}
