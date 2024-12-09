package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName  string `mapstructure:"APP_NAME"`
	Port     string `mapstructure:"PORT"`
	RabbitMQ struct {
		URL string `mapstructure:"RABBITMQ_URL"`
	} `mapstructure:"RABBITMQ"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	log.Println("Configuration loaded successfully")
	return &config, nil
}
