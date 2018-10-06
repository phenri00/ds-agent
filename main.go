package main

import (
	"log"
	"net/http"
)

func main() {

	envs := getEnv()
	http.HandleFunc("/containers", middleWareAuth(envs.Secret, envs.listContainers))
	http.HandleFunc("/services", middleWareAuth(envs.Secret, envs.listServices))
	http.HandleFunc("/services/update", middleWareAuth(envs.Secret, envs.updateService))
	log.Print("INFO: Server listning at port: ", envs.Port)

	switch envs.Tls {
	case true:
		log.Print("INFO: TLS enabled")
		log.Fatal(http.ListenAndServeTLS(":"+envs.Port, "./cert.pem", "./key.pem", nil))
	default:
		log.Print("WARNING: TLS disabled")
		log.Fatal(http.ListenAndServe(":"+envs.Port, nil))
	}
}
