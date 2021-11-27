package request

import "investaBackend/business/users"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *UserLogin) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    user.Email,
		Password: user.Password,
	}
}

func (user *UserRegister) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     user.Nama,
		Email:    user.Email,
		Password: user.Password,
	}
}
