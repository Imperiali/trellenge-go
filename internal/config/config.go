package config

import "github.com/spf13/viper"

type Config struct {
	AppName string `mapstructure:"APP_NAME"`
	Port    int    `mapstructure:"PORT"`
}

func New() *Config {
	viper.SetDefault("APP_NAME", "trellenge-go")
	viper.SetDefault("PORT", 3333)
	viper.AutomaticEnv()

	config := &Config{}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	return config
}
