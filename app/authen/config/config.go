package config

import (
	"fmt"
	"os"
	"strconv"
)

// config base

// deploy env.
const (
	DeployEnvDev  = "dev"
	DeployEnvStag = "stag"
	DeployEnvProd = "prod"
)

// Config hold http/grpc server config
type ServerConfig struct {
	Addr string
	Port int
}

// DefaultServerConfig return a default server config
func AppServerConfig() ServerConfig {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		fmt.Printf("Init AppServerConfig error: %v", err)
		return ServerConfig{}
	}
	return ServerConfig{
		Addr: fmt.Sprintf("0.0.0.0:%d", port),
		Port: port,
	}
}

// Config ...
type Base struct {
	Env    string       `json:"env" mapstructure:"env"`
	Server ServerConfig `json:"server" mapstructure:"server"`
}

func (b Base) IsDevelopment() bool {
	return b.Env == DeployEnvDev
}

func LoadAppConfig() *Base {
	return &Base{
		Env:    DeployEnvDev,
		Server: AppServerConfig(),
	}
}
