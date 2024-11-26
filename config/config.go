package config

import "github.com/spf13/viper"

type Config struct {
	Port int `mapstructure:"PORT"`
}

func LoadConfig() (*Config, error) {
	var config Config

	viper.AutomaticEnv()
	viper.BindEnv("PORT")

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
