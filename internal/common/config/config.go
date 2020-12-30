package config

//CF is the config variable and can be used to fetch from Config struct
var CF *Config

//Config holds all config data for service
type Config struct {
	Environment string `json:"env"`
	Server      server `json:"server"`
}

type (
	server struct {
		Port int `json:"port"`
	}
)
