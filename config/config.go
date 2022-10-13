package config

import (
	"github.com/joeshaw/envdecode"
	"time"
)

type Conf struct {
	Server   Server
	Database Database
	Debug    bool `env:"DEBUG,required"`
}

type Server struct {
	Port         int           `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
}

type Database struct {
	Port     int    `env:"DB_PORT,required"`
	Host     string `env:"DB_HOST,required"`
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASS,required"`
	DBName   string `env:"DB_NAME,required"`
}

func NewConfig() (*Conf, error) {
	var config Conf
	if err := envdecode.StrictDecode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
