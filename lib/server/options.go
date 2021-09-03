package server

// Option configures a gRPC and a gateway server.
type Option func(*Config)

func createConfig(opts []Option) *Config {
	c := createDefaultConfig()
	for _, f := range opts {
		f(c)
	}
	return c
}

// WithGrpcAddrListen
func WithGrpcAddrListen(l Listen) Option {
	return func(c *Config) {
		c.Grpc.Addr = l
	}
}

// WithServiceServer
func WithServiceServer(srv ...ServiceServer) Option {
	return func(c *Config) {
		c.ServiceServers = append(c.ServiceServers, srv...)
	}
}
