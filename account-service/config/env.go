package config

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	envVarNotDefinedError = "%s %s is not defined as env variable"
)

type Env struct {
	DatabaseName   string
	MongoUri       string
	EncryptsionKey string
}

func (env *Env) GetDbName() string {
	return env.DatabaseName
}

func (env *Env) GetDbUri() string {
	return env.MongoUri
}

func NewEnv() Env {
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

	encryptsionKey, err := GetEnvVar("ENCRYPTION_TOKEN")
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		log.Fatal(strings.Join(errs, " "))
	}

	env.DatabaseName = databaseName
	env.MongoUri = mongoUri
	env.EncryptsionKey = encryptsionKey

	return env
}

func GetEnvVar(key string) (string, error) {
	res := os.Getenv(key)
	if res == "" {
		return "", fmt.Errorf(envVarNotDefinedError, "[ERROR]", key)
	}

	return res, nil
}
