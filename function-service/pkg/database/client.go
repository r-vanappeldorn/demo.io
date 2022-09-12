package database

import (
	"function-srv/config"
	"log"
)

func ConnectToDb(env *config.Env) DatabaseHelper {
	client, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect()
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(env.DatabaseName)
}