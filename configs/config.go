package configs

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"os"
)

var FilePath = ""

var Config Configuration

func init() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Config)
	if err != nil {
		log.Fatal("CANT LOGIN CONFIG")
	}
}

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
