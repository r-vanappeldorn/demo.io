package main

import (
	"fmt"
	"function-srv/cmd/function-service/app"
	"function-srv/config"
	"function-srv/pkg/database"
	"function-srv/pkg/messages"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

func main() {
	env := config.GetEnv()
	cfg := config.New(env)

	logger := log.New(
		os.Stdout,
		color.New(color.FgGreen).Sprintf("[%s] ", cfg.SrvName),
		log.Ldate|log.Ltime,
	)

	db := database.ConnectToDb(&env)
	app := app.NewApplication(logger, cfg, db)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.Router(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("%s starting %s service at %s", messages.Succes(), cfg.Env, srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
