package auth



type (
	RegisterRequest struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"email,required"`
		Password string `json:"password"`
	}
	RegisterResponse struct {
		Token string     `json:"token"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"email,required"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		Token string     `json:"token"`
	}
)
