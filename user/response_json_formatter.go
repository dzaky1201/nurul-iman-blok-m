package user

import (
	"nurul-iman-blok-m/model"
)

type UserFormatter struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

func UserJsonFormatter(user model.User, role string, token string) UserFormatter {

	formatter := UserFormatter{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  role,
		Token: token,
	}

	return formatter
}
