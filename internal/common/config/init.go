package config

import (
	"fmt"
	"log"

	"github.com/baghelrahul159/deadline-tracker/internal/model"
	"github.com/spf13/viper"
)

//Load loads variables from config file
func Load() {
	CF = &Config{}
	viper.AutomaticEnv()
	env := viper.GetString("APP_ENV")
	if env == "" {
		env = model.EnvDevelopment
	}
	viper.SetConfigName("tracker." + env)
	viper.SetConfigType("json")
	viper.AddConfigPath("files/etc")
	viper.AddConfigPath("etc")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	err = viper.Unmarshal(CF)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

//GetConfig returns config variable which can then be used for fetching from config struct
func GetConfig() *Config {
	return CF
}
