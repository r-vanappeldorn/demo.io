package main

import (
	"account-service/cmd/account-service/app"
	"account-service/config"
	"account-service/pkg/database"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	env := config.NewEnv()
	cnf := config.New(env)

	db := database.ConnectToDb(&env)

	if err := app.New(logger, cnf, db).Engine.Run(":3000"); err != nil {
		logger.Fatal(err)
	}
}
