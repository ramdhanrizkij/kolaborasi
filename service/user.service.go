package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"kolaborasi/dto"
	"kolaborasi/entity"
	"kolaborasi/repository"
)

type UserService interface {
	RegisterUser(input dto.RegisterUserDto) (entity.User, error)
	LoginUser(input dto.LoginDTO) (entity.User, error)
	IsEmailAvailable(input dto.EmailCheckDTO) (bool, error)
	SaveAvatar(ID int, fileLocation string) (entity.User, error)
	GetUserById(ID int) (entity.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) RegisterUser(input dto.RegisterUserDto) (entity.User, error) {
	user := entity.User{}
	pass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		panic("gagal generate password")
	}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	user.PasswordHash = string(pass)
	user.Role = "user"

	newUser, err := s.repo.InsertUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *userService) LoginUser(input dto.LoginDTO) (entity.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) IsEmailAvailable(input dto.EmailCheckDTO) (bool, error) {
	email := input.Email
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return false, err
	}
	fmt.Printf("%v", user.ID)
	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *userService) SaveAvatar(ID int, fileLocation string) (entity.User, error) {
	user, err := s.repo.FindUserById(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation
	updateUser, err := s.repo.UpdateUser(user)
	if err != nil {
		return user, err
	}

	return updateUser, nil
}

func (s *userService) GetUserById(ID int) (entity.User, error) {
	user, err := s.repo.FindUserById(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that ID")
	}

	return user, nil
}
