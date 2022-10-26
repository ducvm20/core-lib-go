package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Keys  map[string]interface{} `mapstructure:"keys"`
	Urls  map[string]interface{} `mapstructure:"urls"`
	Eps   map[string]interface{} `mapstructure:"eps"`
	Ports map[string]interface{} `mapstructure:"ports"`
}

var Cf *Config
var IsProduction bool = false

// Please contract admin-Ducvm20 if there was changes in stt.json
func LoadConfig(path string) {

	viper.AddConfigPath(path)
	viper.SetConfigName("stt")
	viper.SetConfigType("json")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error loading config", err)
		return
	}
	env := "stag"
	if IsProduction {
		env = "prod"
	}
	Cf = &Config{
		Ports: viper.Get("ports").(map[string]interface{}),
		Eps:   viper.Get("eps").(map[string]interface{}),
		Keys:  viper.Get(fmt.Sprintf("keys.%s", env)).(map[string]interface{}),
		Urls:  viper.Get(fmt.Sprintf("urls.%s", env)).(map[string]interface{}),
	}
}
func Cfg(key string) interface{} {
	idx := strings.Index(key, ".")
	fmt.Println("Config not found", key)
	if idx != -1 {
		if strings.Contains(key, "keys") {
			return Cf.Keys[key[idx+1:]]
		}
		if strings.Contains(key, "eps") {
			return Cf.Eps[key[idx+1:]]
		}
		if strings.Contains(key, "ports") {
			return Cf.Ports[key[idx+1:]]
		}
		if strings.Contains(key, "urls") {
			return Cf.Urls[key[idx+1:]]
		}
	} else {
		return viper.Get(key)
	}
	return ""
}
