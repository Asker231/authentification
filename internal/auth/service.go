package auth

import (
	"errors"
	"fmt"

	"github.com/Asker231/authentification.git/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{
	repo *user.UserRepoSitory
}

func NewAuthService(repo *user.UserRepoSitory)*AuthService{
	return &AuthService{
		repo: repo,
	}
}
func(a *AuthService)RegisterUser(email,password,name string)(*user.User,error){
	 userFind := a.repo.FindByEmail(email)
	 if userFind != nil{
		return nil, errors.New("Пользователь уже существует")
	 }
	 passwordHash,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	 if err != nil{
		fmt.Println(err.Error())
	 }
	 user := &user.User{
		Name: name,
		Email: email,
		Password: string(passwordHash),
	 }
	 uc,err := a.repo.CreateUser(user)
	 if err != nil{
		fmt.Println(err.Error())
	 }
	 return uc,nil

}

func(a *AuthService)DeleteUserByID(id int){
	err := a.repo.DeleteUserById(id)
	if err != nil{
		fmt.Println(err.Error())
	}
}