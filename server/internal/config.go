package internal

import "fmt"

// Services represents the configuration for all services
type Services struct {
	DB   DB   `yaml:"db,omitempty"`
	GRPC GRPC `yaml:"gRPC,omitempty"`
	API  API  `yaml:"api,omitempty"`
}

// DB represents the configuration for this service
type DB struct {
	Host     string `yaml:"host,omitempty"`
	Port     string `yaml:"port,omitempty"`
	Dialect  string `yaml:"dialect,omitempty"`
	User     string `yaml:"user,omitempty"`
	DBName   string `yaml:"db_name,omitempty"`
	Password string `yaml:"password,omitempty"`
}

// Check the struct of missing values
func (s DB) Check() error {
	if s.Host == "" {
		return fmt.Errorf("missing value 'host'")
	}
	if s.Port == "" {
		return fmt.Errorf("missing value 'port'")
	}
	if s.Dialect == "" {
		return fmt.Errorf("missing value 'dialect'")
	}
	if s.User == "" {
		return fmt.Errorf("missing value 'user'")
	}
	if s.DBName == "" {
		return fmt.Errorf("missing value 'db_name'")
	}
	if s.Password == "" {
		return fmt.Errorf("missing value 'password'")
	}
	return nil
}

// GRPC represents the configuration for this service
type GRPC struct {
	Host string `yaml:"host,omitempty"`
	Port string `yaml:"port,omitempty"`
}

// Check the struct of missing values
func (s GRPC) Check() error {
	if s.Host == "" {
		return fmt.Errorf("missing value 'host'")
	}
	if s.Port == "" {
		return fmt.Errorf("missing value 'port'")
	}
	return nil
}

// API represents the configuration for this service
type API struct {
	Host     string `yaml:"host,omitempty"`
	Port     string `yaml:"port,omitempty"`
	GrpcHost string `yaml:"gRPCHost,omitempty"`
	GrpcPort string `yaml:"gRPCPort,omitempty"`
}

// Check the struct of missing values
func (s API) Check() error {
	if s.Host == "" {
		return fmt.Errorf("missing value 'host'")
	}
	if s.Port == "" {
		return fmt.Errorf("missing value 'port'")
	}
	if s.GrpcHost == "" {
		return fmt.Errorf("missing value 'gRPCHost'")
	}
	if s.GrpcPort == "" {
		return fmt.Errorf("missing value 'gRPCPort'")
	}
	return nil
}
