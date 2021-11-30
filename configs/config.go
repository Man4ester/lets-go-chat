package configs

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	ServerPort int      `json:"server_port"`
	DBConfig   DBConfig `json:"bd_config"`
	JWTSecret string     `json:"jwt_secret"`
}

type DBConfig struct {
	DBDialect  string `json:"db_dialect"`
	DBHost     string `json:"db_host"`
	DBPort     int    `json:"db_port"`
	DBName     string `json:"db_name"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
}

func LoadConfig(configFile string) (*Configuration, error) {
	var config Configuration
	file, err := os.Open(configFile)
	if err != nil{
		return &config, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return &config, err
	}
	return &config, nil
}
