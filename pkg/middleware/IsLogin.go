package middleware

import (
	"net/http"

	"github.com/Asker231/authentification.git/config"
	"github.com/Asker231/authentification.git/pkg/jwt"
	"github.com/Asker231/authentification.git/pkg/res"
)



func IsLogin(next http.Handler,conf *config.AppConfig)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.URL.Query().Get("Authorization")
		isEmail,email := jwt.NewJWTInit(conf.SECRET).ParseJWT(token)
		if !isEmail{
			res.Response(w,email,404)
			return
		}
		
		next.ServeHTTP(w,r)
	})
}