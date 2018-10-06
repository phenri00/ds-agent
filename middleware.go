package main

import (
	"log"
	"net/http"
)

func middleWareAuth(secret string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("x-auth") != secret {
			log.Print("ERROR: Unauthorized. Remote addr: ", r.RemoteAddr)
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}
		h(w, r)
		log.Println("Auth OK.")
	}
}
