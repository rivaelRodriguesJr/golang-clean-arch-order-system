package configs

import (
	"errors"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
	MQUser            string `mapstructure:"MQ_USER"`
	MQPassword        string `mapstructure:"MQ_PASSWORD"`
	MQHost            string `mapstructure:"MQ_HOST"`
	MQPort            string `mapstructure:"MQ_PORT"`
}

func LoadConfig(envFile string) (*Config, error) {
	v := viper.New()

	v.SetConfigType("env")
	v.AutomaticEnv()

	// Allow environment variables to override config file values.
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	keys := []string{
		"DB_DRIVER",
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"WEB_SERVER_PORT",
		"GRPC_SERVER_PORT",
		"GRAPHQL_SERVER_PORT",
		"MQ_USER",
		"MQ_PASSWORD",
		"MQ_HOST",
		"MQ_PORT",
	}

	for _, key := range keys {
		if err := v.BindEnv(key); err != nil {
			return nil, err
		}
	}

	// .env is optional: environment variables alone are sufficient.
	if envFile != "" {
		if _, err := os.Stat(envFile); err == nil {
			v.SetConfigFile(envFile)

			if err := v.ReadInConfig(); err != nil {
				return nil, err
			}
		} else if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}

	var cfg Config

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
