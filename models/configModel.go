package models

import (
	"text/template"
)

type Config struct { //Структура содержащая поля конфига для rest api
	Addr     string `toml:"addr"`
	LogLevel string `toml:"log_level"`
}

type TemplateRenderer struct {
	Templates *template.Template
}

func NewConfig() *Config { //Создание структуры
	return &Config{}
}
