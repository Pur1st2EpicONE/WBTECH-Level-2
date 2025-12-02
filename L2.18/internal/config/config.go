package config

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type App struct {
	Logger  Logger
	Server  Server
	Service Service
	Storage Storage
}

type Logger struct {
	LogDir string
	Debug  bool
}

type Server struct {
	Port            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	MaxHeaderBytes  int
	ShutdownTimeout time.Duration
}

type Service struct {
	MaxEventsPerUser int
}

type Storage struct {
	ExpectedUsers    int
	MaxEventsPerUser int
	MaxEventsPerDay  int
}

func Load() (App, error) {

	if err := godotenv.Load(); err != nil {
		return App{}, fmt.Errorf("godotenv: failed to %v", err)
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return App{}, fmt.Errorf("viper: %v", err)
	}

	logger := loggerConfig()
	server := serverConfig()
	service := serviceConfig()
	storage := storageConfig()

	failsafe(&logger, &server, &service, &storage)

	return App{
		Logger:  logger,
		Server:  server,
		Service: service,
		Storage: storage,
	}, nil

}

func loggerConfig() Logger {
	return Logger{
		LogDir: viper.GetString("app.logger.log_directory"),
		Debug:  viper.GetBool("app.logger.debug_mode"),
	}
}

func serverConfig() Server {
	return Server{
		Port:            viper.GetString("server.port"),
		ReadTimeout:     viper.GetDuration("server.read_timeout"),
		WriteTimeout:    viper.GetDuration("server.write_timeout"),
		MaxHeaderBytes:  viper.GetInt("server.max_header_bytes"),
		ShutdownTimeout: viper.GetDuration("server.shutdown_timeout"),
	}
}

func serviceConfig() Service {
	return Service{
		MaxEventsPerUser: viper.GetInt("service.max_events_per_user"),
	}
}

func storageConfig() Storage {
	return Storage{
		ExpectedUsers:   viper.GetInt("storage.expected_users"),
		MaxEventsPerDay: viper.GetInt("storage.max_events_per_day"),
	}
}

func failsafe(logger *Logger, server *Server, service *Service, storage *Storage) {

	if len(viper.AllSettings()) == 0 {

		fmt.Println("config file is empty, switching to default values")

		*logger = Logger{Debug: true}
		*server = Server{Port: "8080", ReadTimeout: 5 * time.Second, WriteTimeout: 10 * time.Second, MaxHeaderBytes: 1048576, ShutdownTimeout: 15 * time.Second}
		*service = Service{MaxEventsPerUser: 100}
		*storage = Storage{ExpectedUsers: 100, MaxEventsPerDay: 100, MaxEventsPerUser: 100}

		return

	}

	if !viper.IsSet("app.logger.debug_mode") {
		fmt.Println("logger.debug_mode missing, switching to default 'true'")
		logger.Debug = true
	}

	if !viper.IsSet("server.port") {
		fmt.Println("server.port missing, switching to default '8080'")
		server.Port = "8080"
	}
	if !viper.IsSet("server.read_timeout") {
		fmt.Println("server.read_timeout missing, switching to default 5s")
		server.ReadTimeout = 5 * time.Second
	}
	if !viper.IsSet("server.write_timeout") {
		fmt.Println("server.write_timeout missing, switching to default 10s")
		server.WriteTimeout = 10 * time.Second
	}
	if !viper.IsSet("server.max_header_bytes") {
		fmt.Println("server.max_header_bytes missing, switching to default 1MB")
		server.MaxHeaderBytes = 1048576
	}
	if !viper.IsSet("server.shutdown_timeout") {
		fmt.Println("server.shutdown_timeout missing, switching to default 15s")
		server.ShutdownTimeout = 15 * time.Second
	}

	if !viper.IsSet("service.max_events_per_user") {
		fmt.Println("service.max_events_per_user missing, switching to default 100")
		service.MaxEventsPerUser = 100
	}

	if !viper.IsSet("storage.expected_users") {
		fmt.Println("storage.expected_users missing, switching to default 100")
		storage.ExpectedUsers = 100
	}
	if !viper.IsSet("storage.max_events_per_day") {
		fmt.Println("storage.max_events_per_day missing, switching to default 100")
		storage.MaxEventsPerDay = 100
	}

}
