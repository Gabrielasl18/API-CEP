package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

var config Config

type Config struct {
	MySQLUser     string `required:"true" envconfig:"MYSQL_USER"`
	MySQLPassword string `required:"true" envconfig:"MYSQL_ROOT_PASSWORD"`
	MySQLHost     string `required:"true" envconfig:"MYSQL_HOSTS"`
	MySQLPort     string `required:"true" envconfig:"MYSQL_PORTA"`
	MySQLDB       string `required:"true" envconfig:"MYSQL_DATABASE"`
}

func GetConfig() (Config, error) {
	if err := envconfig.Process("", &config); err != nil {
		return Config{}, fmt.Errorf("failed to fetch app configs: %w", err)
	}

	return config, nil
}
