package request

import "investaBackend/business/user_investasi"

type UserInvestorLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *UserInvestorLogin) ToDomain() *user_investasi.Domain {
	return &user_investasi.Domain{
		Email:    user.Email,
		Password: user.Password,
	}
}
