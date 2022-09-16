package database

import (
	"log"
)

type Env interface {
	GetDbName() string
	GetDbUri() string
}

func ConnectToDb(env Env) DatabaseHelper {
	client, err := NewClient(env)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect()
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(env.GetDbName())
}
