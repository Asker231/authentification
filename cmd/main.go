package main

import (
	"net/http"

	"github.com/Asker231/authentification.git/internal/auth"
)


func main() {
	router := http.NewServeMux()
	//handler
	auth.NewHandleAuth(router)
	server := http.Server{
		Addr: ":3002",
		Handler: router,
	}

	server.ListenAndServe()
}