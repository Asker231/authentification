package auth

import (
	"fmt"
	"net/http"

	"github.com/Asker231/authentification.git/internal/user"
	"github.com/Asker231/authentification.git/pkg/req"
	"github.com/Asker231/authentification.git/pkg/res"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct{
	userRepo *user.UserRepoSitory
}

func NewHandleAuth(router *http.ServeMux,repo user.UserRepoSitory){
	ah := AuthHandler{
		userRepo: &repo,
	}

	router.HandleFunc("POST /auth/register",ah.Register())
	router.HandleFunc("POST /auth/login",ah.Login())
}

func( a * AuthHandler)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		payload , err := req.HandleBody[RegisterRequest](w,r)
		if err != nil{
			res.Response(w,err.Error(),404)
			return
		}
		usr,err := a.userRepo.CreateUser(&user.User{
			Email: payload.Email,
			Name: payload.Name,
			Password: payload.Password,
		})
		if err != nil{
			fmt.Println(err.Error())
		}
		resp := RegisterResponse{
			Token: GenerateToken(),
			User: usr,
		}
		res.Response(w,resp,401)
	}
}



func GenerateToken()string{
	k := []byte("199823231887As")
	t := jwt.New(jwt.SigningMethodHS256)
	s,_:= t.SignedString(k)
	return s
}


func( a * AuthHandler)Login()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {}
}