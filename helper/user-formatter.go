package helper

import (
	"github.com/go-playground/validator/v10"
	"kolaborasi/entity"
)

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

type ProfileFormatter struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Occupation     string `json:"occupation"`
	Email          string `json:"email"`
	AvatarFileName string `json:"avatar"`
	Role           string `json:"role"`
}

func FormatUser(user entity.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}
	return formatter
}

func FormatProfile(user entity.User) ProfileFormatter {
	formatter := ProfileFormatter{
		ID:             user.ID,
		Name:           user.Name,
		Occupation:     user.Occupation,
		Email:          user.Email,
		Role:           user.Role,
		AvatarFileName: user.AvatarFileName,
	}
	return formatter
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
