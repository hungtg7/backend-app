package server

import (
	"fmt"
	"net"
)

// Address represents a network end point address.
type Listen struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
}

func (l Listen) String() string {
	return fmt.Sprintf("%s:%d", l.Host, l.Port)
}

func createDefaultConfig() *Config {
	config := &Config{
		Grpc: createDefaultGrpcConfig(),
		// TODO - add Gateway (HTTP server) later when in need
	}

	return config
}

func (l *Listen) CreateListener() (net.Listener, error) {
	lis, err := net.Listen("tcp", l.String())
	if err != nil {
		return nil, fmt.Errorf("failed to listen %s: %w", l.String(), err)
	}
	return lis, nil
}

type Config struct {
	Grpc           *grpcConfig
	ServiceServers []ServiceServer
}
