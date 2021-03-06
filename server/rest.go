package server

import (
	"os"

	"github.com/Vitokz/Moysklad/handler"
	"github.com/Vitokz/Moysklad/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

type Rest struct { //Структура моего сервера
	Config  *models.Config //конфиг сервера
	Logger  *logrus.Logger //логер сервера
	Router  *echo.Echo  //сам echo сервер
	Handler *handler.Handler //Хэндлер
	Token   *models.Token//токен авторизации
}

func New(handler *handler.Handler, conf *models.Config) *Rest { //Создание настроек моего сервера
	return &Rest{
		Config:  conf,
		Logger:  logrus.New(),
		Router:  echo.New(),
		Handler: handler,
		Token:   &models.Token{},
	}
}

func (r *Rest) getRoutes() { //Ф-ция эндпоинтов
	r.Router.Use(middleware.CORS())
	r.Router.GET("/", r.GetTask)
	r.Router.GET("/auth", r.Auth)
	r.Router.GET("/sort", r.AddDescription)
	r.Router.POST("/makeSupply", r.MakeSupply)
	r.Router.POST("/addOrRefactor", r.addOrRefactor)
	//Добавить эндпоинты получения контр агентов,Добавить получение складов, добавить получение юр лица
}

func (r *Rest) Start() { //Запуск сервера
	err := r.configureLogger()
	if err != nil {
		r.Logger.WithError(err).Error()
		os.Exit(1)
	}
	r.getRoutes()
	r.Router.Logger.Fatal(r.Router.Start(r.Config.Addr))
}

func (r *Rest) configureLogger() error { //Настройка Логгера
	level, err := logrus.ParseLevel(r.Config.LogLevel)
	if err != nil {
		return err
	}
	r.Logger.SetLevel(level)
	return nil
}
