package config

import "flag"

type Config struct {
	Port    int
	Env     string
	SrvName string
	EnvVars Env
}

// creates a new instance of Config
// with the default values
// Port: 3000
// Env: development
// SrvName: function-service
func New(env Env) Config {
	var conf Config
	conf.EnvVars = env

	flag.IntVar(&conf.Port, "port", 3000, "service port")
	flag.StringVar(&conf.Env, "env", "development", "Env (development|stating|production)")
	flag.StringVar(&conf.SrvName, "srvName", "function-service", "service name")
	flag.Parse()

	return conf
}
