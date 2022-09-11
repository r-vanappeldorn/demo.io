package app

import (
	"log"

	"function-srv/config"
)

type Application struct {
	Logger *log.Logger
	Config config.Config
}

// returns a new application
func NewApplication(logger *log.Logger, config config.Config) *Application {
	return &Application{Logger: logger, Config: config}
}
