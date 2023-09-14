package main

import (
	"fmt"
	"github.com/spf13/pflag"
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

	var Role string
	var Host string
	var Port string

	pflag.StringVarP(&Role, "role", "r", "admin", "角色")
	pflag.StringVarP(&Port, "port", "p", "8899", "端口")
	pflag.StringVarP(&Host, "host", "h", "127.0.0.1", "主机，非 admin 角色需要指定 host")
	pflag.Parse()
	fmt.Println(Role, Port, Host)

	viper.SetDefault("Role", Role)
	viper.SetDefault("Port", Port)
	viper.SetDefault("Host", Host)

	_ = viper.ReadInConfig()

	err := viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("fatal to unmarshal config: %w", err))
	}

	fmt.Printf("config %#v\n", config)
}
