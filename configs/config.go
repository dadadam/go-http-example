package configs

import (
	log "github.com/sirupsen/logrus"

	"github.com/caarlos0/env/v8"
)

type Config struct {
	Port       int    `env:"PORT" envDefault:"8080"`
	LogLevel   string `env:"LOG_LEVEL" envDefault:"INFO"`
	DbHost     string `env:"DB_HOST" envDefault:"localhost"`
	DbPort     int    `env:"DB_PORT" envDefault:"5432"`
	DbName     string `env:"DB_NAME" envDefault:"app_db"`
	DbUser     string `env:"DB_USER" envDefault:"app_user"`
	DbPassword string `env:"DB_PASSWORD" envDefault:""`
}

var cfg *Config

func Init() {
	cfg = &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Error(err)
	}

	configureLogger(cfg.LogLevel)
}

func GetConfig() *Config {
	return cfg
}

func configureLogger(logLevel string) {
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(level)
	log.SetFormatter(&log.JSONFormatter{})
}
