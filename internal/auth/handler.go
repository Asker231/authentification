package auth

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Asker231/authentification.git/config"
	"github.com/Asker231/authentification.git/pkg/jwt"
	"github.com/Asker231/authentification.git/pkg/middleware"
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
	router.Handle("POST /auth/login",  middleware.IsLogin( ah.Login(),conf))
	router.HandleFunc("DELETE /auth/delete/{id}", ah.DeleteUser())
	router.HandleFunc("/home",ah.HomePage())

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

		//Чтение body
		payload, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			res.Response(w, err.Error(), 404)
			return
		}
		// Логика регистрации черес методы сервиса auth
		u , err := a.userService.RegisterUser(payload.Email, payload.Password, payload.Name)
		if err != nil {
			fmt.Println(err.Error())
			res.Response(w, err.Error(), 404)
			return
		}

		// Сосздание jwt инстанса
		result,err := jwt.NewJWTInit(a.conf.SECRET).CreateJWT(jwt.JWTData{
			Email: u.Email,
		})
		if err != nil{
			fmt.Println(err.Error())
			return
		}

		// Response JWT token
		res.Response(w,RegisterResponse{
			Token: result,
		},201)
	}
}


func (a *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Чтение body
		body, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// Логика логина черес методы сервиса auth
		u ,err := a.userService.Login(body.Email, body.Password)
		if err != nil {
			res.Response(w, err.Error(), 404)
			return
		}

		t ,err := jwt.NewJWTInit(a.conf.SECRET).CreateJWT(jwt.JWTData{
			Email: u.Email,
		})

		if err != nil{
			fmt.Println(err.Error())
			return
		}
		c := http.Cookie{
			Name: "jwt",
			Value: t,
			Path: "/",
			MaxAge: 3600,
			HttpOnly: true,
			Secure: true,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(w,&c)
		
		http.Redirect(w,r,"/home",301)

		
		
	}
}

func(a *AuthHandler)HomePage()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"static/index.html")	
	}
}
