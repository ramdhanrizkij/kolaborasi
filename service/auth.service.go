package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	GenerateToken(UserID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
type jwtService struct {
}

func NewAuthService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("KOLABORASISECRETJWT")

func (s *jwtService) GenerateToken(UserID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = UserID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signInToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signInToken, err
	}
	return signInToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return token, err
	}

	return token, nil
}
