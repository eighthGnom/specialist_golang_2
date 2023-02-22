package main

import (
	"flag"
	"log"
	"my_standart_web_server/internal/app/api"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "Path to .toml config file")
}

func main() {
	flag.Parse()
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("Cant find config file, used default", err)
	}
	server := api.New(config)
	log.Fatal(server.Start())
}
