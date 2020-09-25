package configs

import (
	"encoding/json"
	"log"
	"os"
)

// Config struct
type Config struct {
	Environment string
	Database    struct {
		ConnectionString string `json:"connectionString"`
		DatabaseName     string `json:"databaseName"`
	} `json:"database"`
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
}

// GlobalConfig is config object
var GlobalConfig *Config

func Setup() {
	environment := os.Getenv("ENVIRONMENT")

	if len(environment) <= 0 {
		environment = "Development"
	}

	log.Printf("Running environment: %s", environment)

	config, err := loadConfiguration(environment)
	if err != nil {
		log.Printf("configs.Setup, fail to parse ", err)
		os.Exit(0)
	}

	log.Printf("Configs: %s", config)

	GlobalConfig = config
}

func loadConfiguration(environment string) (*Config, error) {
	file := "configs/config." + environment + ".json"

	var config *Config

	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	config.Environment = environment

	return config, err
}
