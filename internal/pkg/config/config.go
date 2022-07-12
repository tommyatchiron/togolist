package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Http struct {
		ListenAddr string `mapstructure:"listen_addr"`
	}
	Grpc struct {
		ListenAddr string `mapstructure:"listen_addr"`
	}
	Logs struct {
		Path string `mapstructure:"path"`
	}
	Postgres struct {
		Dsn string `mapstructure:"dsn"`
	}
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")

	viper.AddConfigPath("/etc/togolist")
	viper.AddConfigPath("$XDG_CONFIG_HOME/togolist")
	viper.AddConfigPath("./config")

	// support reading from environmental variables
	// all env variables are capitalized, dot (levels) and dashes are replaced with underscores
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// default variables
	viper.SetDefault("http.listen_addr", ":8080")
	viper.SetDefault("grpc.listen_addr", ":8080")
	viper.SetDefault("logs.path", "/var/log/togolist")

	err := viper.ReadInConfig()

	if err != nil {
		return nil, fmt.Errorf("fail to read config: %w", err)
	}

	var config Config
	err = viper.Unmarshal(&config)

	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal config: %w", err)
	}

	return &config, nil
}
