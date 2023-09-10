package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Role string
	Port string
	Host string
}

var config Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./")

	viper.SetDefault("Role", "admin")
	viper.SetDefault("Port", "8899")
	viper.SetDefault("Host", "127.0.0.1")

	_ = viper.ReadInConfig()

	err := viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("fatal to unmarshal config: %w", err))
	}

	fmt.Printf("config %#v\n", config)
}
