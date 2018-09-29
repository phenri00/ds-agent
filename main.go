package main

import (
	"log"
	"net/http"
)

func main() {

	envs := getEnv()
	http.HandleFunc("/services", listServices)
	http.HandleFunc("/update", envs.updateService)
	log.Print("Server listning at port: ", envs.Port)
	log.Fatal(http.ListenAndServe(":"+envs.Port, nil))
}
