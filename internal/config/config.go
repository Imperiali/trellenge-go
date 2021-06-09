package config

import "github.com/spf13/viper"

type Config struct {
	AppName       string `mapstructure:"APP_NAME"`
	Port          int    `mapstructure:"PORT"`
	CacheAddress  string `mapstructure:"CACHE_ADDRESS"`
	CachePassword string `mapstructure:"CACHE_PASSWORD"`
}

func New() *Config {
	viper.SetDefault("APP_NAME", "trellenge-go")
	viper.SetDefault("PORT", 3333)
	viper.SetDefault("CACHE_ADDRESS", "localhost:6379")
	viper.SetDefault("CACHE_PASSWORD", "")
	viper.AutomaticEnv()

	config := &Config{}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	return config
}
