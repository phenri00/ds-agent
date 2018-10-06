package main

import (
	"log"
	"net/http"
)

func middleWareAuth(secret string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Print(r.RemoteAddr, " ", r.Method, " ", r.URL)

		if r.Header.Get("x-auth") != secret {
			log.Print("ERROR: Unauthorized")
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}
		h(w, r)
		log.Println("Auth OK.")
	}
}
