package configs

import (
	"encoding/json"
	"os"

	"github.com/labstack/gommon/log"
)

type Configuration struct {
	ServerPort int      `json:"server_port"`
	DBConfig   DBConfig `json:"bd_config"`
}

type DBConfig struct {
	DBDialect  string `json:"db_dialect"`
	DBHost     string `json:"db_host"`
	DBPort     int    `json:"db_port"`
	DBName     string `json:"db_name"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
}

func LoadConfig(configFile string) *Configuration {
	var config Configuration
	file, _ := os.Open(configFile)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)
	if err != nil {
		log.Fatal("cant load config")
	}
	return &config
}
