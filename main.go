package main

import (
	"log"
	"net/http"
)

func main() {

	envs := getEnv()
	http.HandleFunc("/services", listServices)
	http.HandleFunc("/services/update", envs.updateService)
	log.Print("Server listning at port: ", envs.Port)
	log.Fatal(http.ListenAndServeTLS(":"+envs.Port, "./cert.pem", "./key.pem", nil))
}
