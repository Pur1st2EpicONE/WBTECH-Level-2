package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/beevik/ntp"
	"github.com/spf13/viper"
)

const (
	configError = 1
	serverError = 2
)

type Config struct {
	Server          string `mapstructure:"server"`
	ShowAll         bool   `mapstructure:"show_all"`
	ShowDate        bool   `mapstructure:"show_date"`
	ShowPreciseTime bool   `mapstructure:"show_precise_time"`
	ShowOffset      bool   `mapstructure:"show_offset"`
	ShowTimeZone    bool   `mapstructure:"show_time_zone"`
}

func main() {

	config, err := loadConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(configError)
	}

	time, err := ntp.Time(config.Server)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(serverError)
	}

	if config.ShowAll {
		fmt.Println(time)
	} else {
		fmt.Println(format(time, config))
	}

}

// loadConfig reads and parses the configuration file
func loadConfig() (Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("failed to load config: %v", err)
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config: %v", err)
	}
	return cfg, nil
}

// format formats the time according to the configuration
func format(time time.Time, config Config) string {
	var str strings.Builder
	if config.ShowDate {
		str.WriteString(time.Format("2006-01-02 "))
	}
	if config.ShowPreciseTime {
		str.WriteString(time.Format("15:04:05.999999999 "))
	} else {
		str.WriteString(time.Format("15:04:05 "))
	}
	if config.ShowOffset {
		str.WriteString(time.Format("-0700 "))
	}
	if config.ShowTimeZone {
		str.WriteString(time.Format("MST"))
	}
	return str.String()
}
