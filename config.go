package main

import (
	"log"
	"os"
)

type Configuration struct {
	Port             string
	RegistryUser     string
	RegistryPassword string
}

func getEnv() Configuration {

	configuration := Configuration{}

	configuration.Port = checkEnv("DS_AGENT_PORT")
	configuration.RegistryUser = checkEnv("DS_AGENT_REGISTRY_USERNAME")
	configuration.RegistryPassword = checkEnv("DS_AGENT_REGISTRY_PASSWORD")

	return configuration
}

func checkEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal("Missing enviroment variable: ", key)
	}
	return value
}
