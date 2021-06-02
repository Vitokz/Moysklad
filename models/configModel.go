package models

type Config struct { //Структура содержащая поля конфига для rest api
	Addr     string `toml:"addr"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config { //Создание структуры
	return &Config{}
}
