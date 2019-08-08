package server

// GrpcConf specifies a gRPC server configuration.
type GrpcConf struct {
	Bin         string `mapstructure:"bin"`
	Register    string `mapstructure:"register"`
	Implementor string `mapstructure:"implementor"`
}

// GatewayConf specifies gRPC gateway server configuration.
type GatewayConf struct {
	Bin      string `mapstructure:"bin"`
	Register string `mapstructure:"register"`
}

// Service is a single service consisting on many servers.
type Service struct {
	Version string      `mapstructure:"version"`
	GRPC    GrpcConf    `mapstructure:"grpc"`
	Gateway GatewayConf `mapstructure:"gateway"`
}
