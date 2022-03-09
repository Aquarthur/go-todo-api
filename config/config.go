package config

import (
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port   int
	Health string
}

type PostgresConfig struct {
	Host     string
	Database string
	Username string
	Password string
}

type Config struct {
	Server   *ServerConfig   `yaml:"server"`
	Postgres *PostgresConfig `yaml:"db"`
}

func NewConfig() (*Config, error) {
	if err := loadConfigFile(); err != nil {
		return nil, err
	}
	config := new(Config)
	return config, viper.Unmarshal(config)
}

func loadConfigFile() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}
