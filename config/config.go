package config

import "log/slog"

type Config struct {
	Host string
	Port string
}

func DefaultConfig() Config {
	slog.Info("config created")
	return Config{
		Host: "localhost",
		Port: "8080",
	}
}
