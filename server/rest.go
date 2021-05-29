package server

import (
	"os"

	"github.com/Vitokz/Moysklad/handler"
	"github.com/Vitokz/Moysklad/models"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type Rest struct {
	Config  *models.Config
	Logger  *logrus.Logger
	Router  *echo.Echo
	Handler *handler.Handler
}

func New(handler *handler.Handler, conf *models.Config) *Rest {
	return &Rest{
		Config:  conf,
		Logger:  logrus.New(),
		Router:  echo.New(),
		Handler: handler,
	}
}

func (r *Rest) getRoutes() {
	r.Router.GET("/", r.GetTask)
	r.Router.GET("/auth", r.Auth)
}

func (r *Rest) Start() {
	err := r.configureLogger()
	if err != nil {
		r.Logger.WithError(err).Error()
		os.Exit(1)
	}
	r.getRoutes()

	r.Router.Logger.Fatal(r.Router.Start(r.Config.Addr))
}

func (r *Rest) configureLogger() error {
	level, err := logrus.ParseLevel(r.Config.LogLevel)
	if err != nil {
		return err
	}
	r.Logger.SetLevel(level)
	return nil
}
