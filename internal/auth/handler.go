package auth

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Asker231/authentification.git/config"
	"github.com/Asker231/authentification.git/pkg/jwt"
	"github.com/Asker231/authentification.git/pkg/req"
	"github.com/Asker231/authentification.git/pkg/res"
)

type AuthHandler struct {
	userService *AuthService
	conf *config.AppConfig
}

func NewHandleAuth(router *http.ServeMux, service *AuthService,conf *config.AppConfig) {
	ah := AuthHandler{
		userService: service,
		conf: conf,
	}

	router.HandleFunc("POST /auth/register", ah.Register())
	router.HandleFunc("POST /auth/login", ah.Login())
	router.HandleFunc("GET /auth/home", ah.Home())
	router.HandleFunc("DELETE /auth/delete/{id}", ah.DeleteUser())

}

func (a *AuthHandler) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println(err.Error())
		}
		err = a.userService.DeletedUserByID(id)
		if err != nil{
			fmt.Println(err.Error())
		}
	}
}

func (a *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			res.Response(w, err.Error(), 404)
			return
		}

		u , err := a.userService.RegisterUser(payload.Email, payload.Password, payload.Name)
		if err != nil {
			fmt.Println(err.Error())
			res.Response(w, err.Error(), 404)
			return
		}
		result,err := jwt.NewJWTInit(a.conf.SECRET).CreateJWT(u.Email)
		if err != nil{
			fmt.Println(err.Error())
			return
		}
		http.Redirect(w, r, "/auth/home", 301)
		res.Response(w,RegisterResponse{
			Token: result,
		},201)
	}
}



func (a *AuthHandler) Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home Page"))
	}
}

func (a *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		u,err := a.userService.Login(body.Email, body.Password)
		if err != nil {
			res.Response(w, err.Error(), 404)
			return
		}
		result,err := jwt.NewJWTInit(a.conf.SECRET).CreateJWT(u.Email)
		if err != nil{
			fmt.Println(err.Error())
			return
		}

		res.Response(w,RegisterResponse{
			Token: result,
		},201)

		http.Redirect(w, r, "/auth/home", 301)
	}
}
