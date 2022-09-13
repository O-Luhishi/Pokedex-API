package config

import (
	"time"
)

type Conf struct {
	Server serverConfig
	Debug  bool
}

type serverConfig struct {
	Port         int
	TimeoutRead  time.Duration
	TimeoutWrite time.Duration
	TimeoutIdle  time.Duration
}

// NewConfig Todo: We'll set these an env variables later when we dockerize the application!
func NewConfig() *Conf {
	return &Conf{
		Server: serverConfig{
			Port:         8080,
			TimeoutRead:  5,
			TimeoutWrite: 180,
			TimeoutIdle:  15,
		},
		Debug: false,
	}
}
