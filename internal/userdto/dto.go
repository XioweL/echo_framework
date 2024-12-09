package internal

type LoginRequest struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
}
type Users struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
