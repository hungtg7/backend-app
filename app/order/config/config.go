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
	Add string
}

// DefaultServerConfig return a default server config
func AppServerConfig() ServerConfig {
	return ServerConfig{
		Add: "0.0.0.0:8888",
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
