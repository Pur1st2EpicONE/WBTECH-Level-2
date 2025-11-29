package config

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type App struct {
	Server  Server
	Logger  Logger
	Storage Storage
	Service Service
}

type Server struct {
	Port            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	MaxHeaderBytes  int
	ShutdownTimeout time.Duration
}

type Storage struct {
	MaxUsers         int
	MaxEventsPerUser int
	MaxEventsPerDay  int
}

type Service struct {
	MaxEventsPerUser int
}

type Logger struct {
	LogDir string
	Debug  bool
}

func Load() (App, error) {
	if err := godotenv.Load(); err != nil {
		return App{}, fmt.Errorf("godotenv — failed to %v", err)
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return App{}, fmt.Errorf("viper — %v", err)
	}

	return App{
		Server:  srvConfig(),
		Logger:  loggerConfig(),
		Storage: storageConfig(),
	}, nil
}

func srvConfig() Server {
	return Server{
		Port:            viper.GetString("server.port"),
		ReadTimeout:     viper.GetDuration("server.read_timeout"),
		WriteTimeout:    viper.GetDuration("server.write_timeout"),
		MaxHeaderBytes:  viper.GetInt("server.max_header_bytes"),
		ShutdownTimeout: viper.GetDuration("server.shutdown_timeout"),
	}
}

func storageConfig() Storage {
	return Storage{
		MaxUsers:         viper.GetInt("storage.max_users"),
		MaxEventsPerUser: viper.GetInt("storage.max_events_per_user"),
		MaxEventsPerDay:  viper.GetInt("storage.max_events_per_day"),
	}
}

func loggerConfig() Logger {
	return Logger{
		LogDir: viper.GetString("app.logger.log_directory"),
		Debug:  viper.GetBool("app.logger.debug_mode"),
	}
}
