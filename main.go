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

	switch envs.Tls {
	case true:
		log.Print("INFO: TLS enabled")
		log.Fatal(http.ListenAndServeTLS(":"+envs.Port, "./cert.pem", "./key.pem", nil))
	default:
		log.Print("WARNING: TLS disabled")
		log.Fatal(http.ListenAndServe(":"+envs.Port, nil))
	}

	log.Fatal(http.ListenAndServe(":"+envs.Port, nil))
}
