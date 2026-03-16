package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Logger LoggerConfig
	Jaeger JaegerConfig
	Models ModelsConfig
}
type ServerConfig struct {
	BindAddr string `mapstructure:"bind_addr"`
}
type LoggerConfig struct {
	Level string `mapstructure:"level"`
}
type JaegerConfig struct {
	URL string `mapstructure:"url"`
}
type ModelsConfig struct {
	Chat  ModelConfig `mapstructure:"chat"`
	Image ModelConfig `mapstructure:"image"`
	Audio ModelConfig `mapstructure:"audio"`
}
type ModelConfig struct {
	Endpoint string `mapstructure:"endpoint"`
	APIKey   string `mapstructure:"api_key"`
}

func LoadConfig(path string) (*Config, error) {

	viper.SetConfigFile(path)

	viper.SetEnvPrefix("AI_BALANCER")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
