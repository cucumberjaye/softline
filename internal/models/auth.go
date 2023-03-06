package models

type RegisterUser struct {
	Login                string `json:"login" validate:"required,alphanum,gte=3"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,gte=6"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
	PhoneNumber          string `json:"phone_number" validate:"required,e164"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6"`
}
