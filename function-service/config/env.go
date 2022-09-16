package config

import (
	"fmt"
	"function-srv/pkg/messages"
	"log"
	"os"
	"strings"
)

const (
	envVarNotDefinedError = "%s %s is not defined as env variable"
)

type Env struct {
	DatabaseName string
	MongoUri     string
}

func GetEnv() Env {
	var env Env
	var errs []string

	mongoUri, err := GetEnvVar("MONGO_URI")
	if err != nil {
		errs = append(errs, err.Error())
	}

	databaseName, err := GetEnvVar("DATABASE_NAME")
	if err != nil {
		errs = append(errs, err.Error())
	}

	env.DatabaseName = databaseName
	env.MongoUri = mongoUri

	if len(errs) > 0 {
		log.Fatal(strings.Join(errs, " "))
	}

	return env
}

func GetEnvVar(key string) (string, error) {
	res := os.Getenv(key)
	if res == "" {
		return "", fmt.Errorf(envVarNotDefinedError, messages.Error(), key)
	}

	return res, nil
}
