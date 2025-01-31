package auth

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Asker231/authentification.git/pkg/req"
	"github.com/Asker231/authentification.git/pkg/res"
	"github.com/Asker231/authentification.git/pkg/jwt"
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
		result,err := jwt.NewJWTInit("some199823231998").CreateJWT(u.Email)
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
		result,err := jwt.NewJWTInit("9d06069f339fb58f02236c1adca8cbeb87db24e78cfc83d30a1a0226f0226c346c6e5a28c14036c52f289789f3fe109759e1f677596c7d34086299bc332151af58f5dbf97942af77cc9c261e3db7be92cf33b1be7f8df997fd55b51a6b522ec54b514b456fff40f9fad60cb962926dba9d0f6c9a7eb0c8abf4156a4fa697b91f7d5fae966d441840532db37e704ad34b4a68208a9beae9dbd198aadb04eb9884968c2c860aa33a669850e22797e2eb7568ac4f041935b82bb52736620626c4547f57516e1ada720fd0bb1c0665681bdf2fb53a2a129cbfbbe80fee6d2601e302b8f542ffe96ce2e0cc8357dcc6ee1b5b155563783f5275fd7791d8d30c9a8a8d").CreateJWT(u.Email)
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
