package config

// config base

// deploy env.
const (
	DeployEnvDev  = "dev"
	DeployEnvStag = "stag"
	DeployEnvProd = "prod"
)

// Config hold http/grpc server config
type ServerConfig struct {
	Host string
	Port int
}

// DefaultServerConfig return a default server config
func AppServerConfig() ServerConfig {
	return ServerConfig{
		Host: "0.0.0.0",
		Port: 10550,
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
