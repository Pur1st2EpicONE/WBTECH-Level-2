// Package downloader contains utilities for saving files to disk and
// reporting the result of save operations.
package downloader

import (
	"fmt"
	"os"
	"path/filepath"
)

// SavedFile describes the result of an attempt to save a file to disk.
// Path is the destination path, Size is the written byte count, Err is any
// error that occurred during the operation.
type SavedFile struct {
	Path string
	Size int
	Err  error
}

// Save writes the provided file bytes to the given path. Directories are
// created as needed. The returned SavedFile contains the path, number of
// bytes written and any error that occurred.
func Save(path string, file []byte) SavedFile {

	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		return SavedFile{Path: path, Size: 0, Err: fmt.Errorf("unable to create directory: %w", err)}
	}

	if err := os.WriteFile(path, file, 0777); err != nil {
		return SavedFile{Path: path, Size: 0, Err: fmt.Errorf("unable to write: %w", err)}
	}

	return SavedFile{Path: path, Size: len(file), Err: nil}

}
