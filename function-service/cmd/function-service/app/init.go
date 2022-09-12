package app

import (
	"log"

	"function-srv/config"
	"function-srv/pkg/database"
)

type Application struct {
	Logger *log.Logger
	Config config.Config
	Db database.DatabaseHelper
}

// returns a new application
func NewApplication(logger *log.Logger, config config.Config, db database.DatabaseHelper) *Application {
	return &Application{Logger: logger, Config: config, Db: db}
}
