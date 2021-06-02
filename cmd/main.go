package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Vitokz/Moysklad/file"
	"github.com/Vitokz/Moysklad/handler"
	"github.com/Vitokz/Moysklad/models"
	"github.com/Vitokz/Moysklad/proto"
	"github.com/Vitokz/Moysklad/server"
)

func main() {
	config := models.NewConfig()  //Создаю модель конфига 
	_, err := toml.DecodeFile(proto.CONFIG_REST_PATH, config) //Заполняю модель конфига с помощью toml	
	if err != nil {
		log.Fatal()
	}

	file := file.NewFileOpen()  //Открываю xlsx ФАИЛ C с соотношениями из 1с
	handler := handler.NewHandler(file) //Создаю хэндлер и засовываю туда структуру открытого файла

	rest := server.New(handler, config)  //Создаю Rest Api localhost

	rest.Start() //Запускаю свой api
}
