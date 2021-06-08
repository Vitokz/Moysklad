package server

import (
	"os"

	"github.com/Vitokz/Moysklad/handler"
	"github.com/Vitokz/Moysklad/models"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type Rest struct { //Структура моего сервера
	Config  *models.Config
	Logger  *logrus.Logger
	Router  *echo.Echo
	Handler *handler.Handler
	Token   *models.Token
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
	r.Router.GET("/", r.GetTask)
	r.Router.GET("/auth", r.Auth)
	r.Router.GET("/sort", r.AddDescription)
	r.Router.GET("/createPrice", r.CreatePrice)
	r.Router.GET("/makeSupply", r.MakeSupply)
	r.Router.GET("/makeProduct", r.MakeProduct)
	r.Router.GET("/refactorProduct",r.RefactorProduct)
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
