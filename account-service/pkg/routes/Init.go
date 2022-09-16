package routes

import (
	"account-service/config"
	"account-service/pkg/database"
	"log"
)

type Routes struct {
	Logger *log.Logger
	Db     database.DatabaseHelper
	cfg    config.Config
}

func New(logger *log.Logger, db database.DatabaseHelper, cfg config.Config) *Routes {
	return &Routes{logger, db, cfg}
}
