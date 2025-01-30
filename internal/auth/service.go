package auth

import "github.com/Asker231/authentification.git/internal/user"

type AuthService struct{
	repo *user.UserRepoSitory
}

func NewAuthService(repo *user.UserRepoSitory)*AuthService{
	return &AuthService{
		repo: repo,
	}
}
