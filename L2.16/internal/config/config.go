// Package config provides configuration loading and default values for the
// wget-like scraper. It uses viper to read a configuration file named
// "config" from the current working directory and exposes helper errors and
// defaults used across the application.
package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var (
	ErrLoadConfig      = errors.New("unable to load")
	ErrUnmarshalConfig = errors.New("failed to unmarshal config")
	ErrNoSelectors     = errors.New("no selectors specified")
)

// Config holds runtime options for the scraper.
// Fields are tagged for viper mapstructure unmarshaling.
type Config struct {
	Async             bool          `mapstructure:"async"`
	Parallelism       int           `mapstructure:"parallelism"`
	Delay             time.Duration `mapstructure:"delay"`
	RandomDelay       time.Duration `mapstructure:"random_delay"`
	UserAgent         string        `mapstructure:"user_agent"`
	RequestTimeout    time.Duration `mapstructure:"request_timeout"`
	IgnoreRobotsTxt   bool          `mapstructure:"ignore_robots_txt"`
	DownloadHTML      bool          `mapstructure:"download_html"`
	DownloadCSS       bool          `mapstructure:"download_css"`
	DownloadScripts   bool          `mapstructure:"download_scripts"`
	DownloadImages    bool          `mapstructure:"download_images"`
	DownloadVideos    bool          `mapstructure:"download_videos"`
	DownloadAudio     bool          `mapstructure:"download_audio"`
	DownloadIframes   bool          `mapstructure:"download_iframes"`
	DownloadFonts     bool          `mapstructure:"download_fonts"`
	DownloadIcons     bool          `mapstructure:"download_icons"`
	DownloadManifests bool          `mapstructure:"download_manifests"`
	DownloadJSON      bool          `mapstructure:"download_json"`
}

// Load reads configuration from a file named "config" in the current
// directory and unmarshals it into a Config. If the file cannot be read or
// unmarshaled an error is returned. If no download selectors are enabled,
// ErrNoSelectors is returned.
func Load() (*Config, error) {

	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	config := new(Config)

	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("%w: %w", ErrLoadConfig, err)
	}

	if err := viper.Unmarshal(config); err != nil {
		return config, fmt.Errorf("%w: %w", ErrUnmarshalConfig, err)
	}

	if !config.DownloadHTML && !config.DownloadCSS && !config.DownloadScripts &&
		!config.DownloadImages && !config.DownloadVideos && !config.DownloadAudio &&
		!config.DownloadIframes && !config.DownloadFonts && !config.DownloadIcons {
		return config, ErrNoSelectors
	}

	return config, nil

}

// Default returns a Config filled with sensible defaults used when the
// configuration file is missing or invalid.
func Default() *Config {
	return &Config{
		Async:             true,
		Parallelism:       5,
		Delay:             500 * time.Millisecond,
		RandomDelay:       200 * time.Millisecond,
		UserAgent:         "RealPerson/1.0",
		RequestTimeout:    15 * time.Minute,
		IgnoreRobotsTxt:   false,
		DownloadHTML:      true,
		DownloadCSS:       true,
		DownloadScripts:   true,
		DownloadImages:    true,
		DownloadVideos:    true,
		DownloadAudio:     true,
		DownloadIframes:   true,
		DownloadFonts:     true,
		DownloadIcons:     true,
		DownloadManifests: true,
		DownloadJSON:      true,
	}
}
