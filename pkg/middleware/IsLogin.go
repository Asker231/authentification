package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Asker231/authentification.git/config"
	"github.com/Asker231/authentification.git/pkg/jwt"
)



func IsLogin(next http.Handler,conf *config.AppConfig)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("auth")
		result := strings.TrimPrefix(auth,"auth ")
		isEmail,email := jwt.NewJWTInit(conf.SECRET).ParseJWT(result)
		fmt.Println(email)
		fmt.Println(isEmail)
		next.ServeHTTP(w,r)
	})
}