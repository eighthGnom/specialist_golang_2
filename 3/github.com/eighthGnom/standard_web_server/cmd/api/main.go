package main

import (
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/eighthGnom/standard_web_server/internal/app/api"
	"github.com/joho/godotenv"
)

var (
	configPath, configFormat string
)

func init() {
	flag.StringVar(&configFormat, "format", ".toml", "File format of the configs")
	flag.StringVar(&configPath, "path", "configs/api.toml", "Path to config file .toml")
}

func main() {
	flag.Parse()
	config := api.NewConfig()
	switch {
	case configFormat == ".env" && configPath != "":
		if err := godotenv.Load(configPath); err != nil {
			log.Println("No .env files", err)
		}
		config.BindAddr = os.Getenv("bind_addr")
		config.LoggerLevel = os.Getenv("logger_level")
		config.StorageConfig.DatabaseURI = os.Getenv("database_uri")

	case configFormat == ".toml" && configPath != "":
		_, err := toml.DecodeFile(configPath, config)
		if err != nil {
			log.Println(err)
		}
	}
	server := api.New(config)
	log.Fatal(server.Start())
}
