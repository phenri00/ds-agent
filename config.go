package main

import (
	"log"
	"os"
	"strconv"
)

type Configuration struct {
	Port             string
	RegistryUser     string
	RegistryPassword string
	Secret           string
	Tls              bool
}

func getEnv() Configuration {

	configuration := Configuration{}

	configuration.Port = checkEnv("DS_AGENT_PORT")
	configuration.RegistryUser = checkEnv("DS_AGENT_REGISTRY_USERNAME")
	configuration.RegistryPassword = checkEnv("DS_AGENT_REGISTRY_PASSWORD")
	configuration.Secret = checkEnv("DS_AGENT_SECRET")
	configuration.Tls, _ = strconv.ParseBool(os.Getenv("DS_AGENT_TLS")) //optional
	return configuration
}

func checkEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal("Missing enviroment variable: ", key)
	}
	return value
}
