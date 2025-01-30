package main

import (
	"fmt"
	"net/http"

	"github.com/Asker231/authentification.git/config"
	"github.com/Asker231/authentification.git/internal/auth"
	"github.com/Asker231/authentification.git/internal/user"
	"github.com/Asker231/authentification.git/pkg/db"
)


func main() {
	router := http.NewServeMux()
	//config
	cnf := config.NewAppConfig()
	//
	database ,err := db.ConnectDataBase(cnf)
	if err != nil{
		fmt.Println(err.Error())
	}
	//repo
    repo :=  user.NewRepoUser(database)
	//servisec
	service := auth.NewAuthService(repo)
	//handler
	auth.NewHandleAuth(router,service)
	server := http.Server{
		Addr: ":3002",
		Handler: router,
	}

	server.ListenAndServe()
}