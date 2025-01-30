package auth

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Asker231/authentification.git/pkg/req"
	"github.com/Asker231/authentification.git/pkg/res"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	userService *AuthService
}

func NewHandleAuth(router *http.ServeMux, service *AuthService) {
	ah := AuthHandler{
		userService: service,
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
		a.userService.DeleteUserByID(id)
	}
}

func (a *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			res.Response(w, err.Error(), 404)
			return
		}

		_, err = a.userService.RegisterUser(payload.Email, payload.Password, payload.Name)
		if err != nil {
			fmt.Println(err.Error())
			res.Response(w, err.Error(), 404)
			return
		}
		http.Redirect(w, r, "/auth/home", 301)
	}
}

func GenerateToken() string {
	k := []byte("199823231887As")
	t := jwt.New(jwt.SigningMethodHS256)
	s, _ := t.SignedString(k)
	return s
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
		_, err = a.userService.Login(body.Email, body.Password)

		if err != nil {
			res.Response(w, err.Error(), 404)
			return
		}
		http.Redirect(w, r, "/auth/home", 301)

	}
}
