package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Vitokz/Moysklad/handler"
	"github.com/Vitokz/Moysklad/models"
	"github.com/Vitokz/Moysklad/proto"
	"github.com/Vitokz/Moysklad/server"
)

func main() {
	config := models.NewConfig()
	_, err := toml.DecodeFile(proto.CONFIG_REST_PATH, config)
	if err != nil {
		log.Fatal()
	}

	handler := handler.NewHandler()

	rest := server.New(handler, config)

	rest.Start()
}
