package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Token string
	Guild string // Set to enable test discord
}

func Load() (*Config, error) {
	// Setup viper
	setupViper()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	// Unmarshal
	conf := &Config{}
	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil

}

func setupViper() {
	// Read in config file
	file()
	// Read in environment variables
	env()
}

func file() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
}

func env() {
	viper.SetEnvPrefix("DISCORD")
	viper.BindEnv("TOKEN")
	viper.BindEnv("GUILD")
	viper.AutomaticEnv()
}
