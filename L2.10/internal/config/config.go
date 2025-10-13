package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ChunkSize int `mapstructure:"chunk_size"`
	Workers   int `mapstructure:"workers"`
}

func Load() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}
	cfg := new(Config)
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}
	return cfg, nil
}
