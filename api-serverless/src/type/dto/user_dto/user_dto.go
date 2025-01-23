package user_dto

type SignUpDTO struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Name     *string `json:"name,omitempty"`
}

type SignInDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
