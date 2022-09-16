package app

import (
	"account-service/config"
	"account-service/pkg/database"
	"account-service/pkg/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Config config.Config
	Engine *gin.Engine
	Logger *log.Logger
	Routes *routes.Routes
	Db     database.DatabaseHelper
}

func New(logger *log.Logger, cnf config.Config, db database.DatabaseHelper) *Application {
	r := gin.Default()
	route := routes.New(logger,db, cnf )
	app := &Application{
		cnf,
		r,
		logger,
		route,
		db,
	}

	app.Router()
	return app
}
