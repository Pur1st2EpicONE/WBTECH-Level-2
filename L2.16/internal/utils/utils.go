// Package utils contains helper utilities used across the application, for
// human-readable sizes, path localization and standardized output.
package utils

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"L2.16/internal/config"
)

const (
	KB = 1024
	MB = 1024 * 1024
)

// WgetHeader prints a wget-like header for a downloaded response.
func WgetHeader(pageURL string, localPath string, contentType string, length int) {
	fmt.Println(time.Now().Format("--2006-01-02 15:04:05--"), pageURL)
	fmt.Println("Response received")
	fmt.Printf("Length: %d (%s) [%s]\n", length, ToHumanSize(int64(length)), contentType)
	fmt.Printf("Saving to: ‘%s’\n\n", localPath)
}

// toHumanSize converts a byte count to a human readable string
func ToHumanSize(bytes int64) string {
	switch {
	case bytes >= MB:
		return fmt.Sprintf("%.1fM", float64(bytes)/MB)
	case bytes >= KB:
		return fmt.Sprintf("%.0fK", float64(bytes)/KB)
	default:
		return fmt.Sprintf("%dB", bytes)
	}
}

// LocalizePath creates a local filesystem path for the given page URL. It
// converts a URL path to a safe local file path and returns it joined with
// the host as the top-level directory.
func LocalizePath(pageURL *url.URL) (string, error) {

	path := strings.TrimPrefix(pageURL.Path, "/")

	if pageURL.RawQuery != "" {
		return pageURL.Host + "/" + path + "?" + pageURL.RawQuery, nil
	}

	if path == "" || strings.HasSuffix(pageURL.Path, "/") {
		path = filepath.Join(path, "index.html")
	}

	localPath, err := filepath.Localize(path)
	if err != nil {
		return "", fmt.Errorf("unable to localize URL: %w", err)
	}

	return filepath.Join(".", pageURL.Host, localPath), nil

}

// FormatOutput prints a formatted summary of the scraping session,
// including total files, total size, elapsed time, and average speed.
func FormatOutput(totalDownloads int, size int64, duration time.Duration) {
	fmt.Printf("FINISHED --%s--\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("Total wall clock time: %.0fs\n", duration.Seconds())
	fmt.Printf("Downloaded: %d files, %s in %ds (%.0f KB/s)\n", totalDownloads, ToHumanSize(size), int(duration.Seconds()), float64(size)/duration.Seconds()/1024)
}

// BuildSelector builds a CSS selector string for colly based on enabled
// download types in the provided config.
func BuildSelector(cfg *config.Config) string {

	selectors := make([]string, 0, 12)

	if cfg.DownloadHTML {
		selectors = append(selectors, "a[href]")
	}
	if cfg.DownloadCSS {
		selectors = append(selectors, "link[rel=stylesheet][href]")
	}
	if cfg.DownloadFonts {
		selectors = append(selectors, "link[rel=preload][as=font][href]")
	}
	if cfg.DownloadScripts {
		selectors = append(selectors, "script[src]")
	}
	if cfg.DownloadImages {
		selectors = append(selectors, "img[src], picture > img[src], img[data-src]")
	}
	if cfg.DownloadVideos {
		selectors = append(selectors, "video[src], video > source[src], source[src]")
	}
	if cfg.DownloadAudio {
		selectors = append(selectors, "audio[src], audio > source[src]")
	}
	if cfg.DownloadIframes {
		selectors = append(selectors, "iframe[src]")
	}
	if cfg.DownloadIcons {
		selectors = append(selectors, "link[rel~='icon'][href]", "link[rel~='shortcut icon'][href]")
	}
	if cfg.DownloadManifests {
		selectors = append(selectors, "link[rel=manifest][href]")
	}
	if cfg.DownloadJSON {
		selectors = append(selectors, "link[rel=preload][as=fetch][href]", "a[href$='.json']")
	}

	return strings.Join(selectors, ", ")

}

// LogError writes an error message to stderr, including optional context.
// If err is nil, only the context is printed.
func LogError(context string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", context, err)
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", context)
	}
}
