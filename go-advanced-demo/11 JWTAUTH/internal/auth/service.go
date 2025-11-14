package auth

import (
	"11-JWTAUTH/internal/user"
	"errors"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service AuthService) Register(email, password, name string) (string, error) {
	existsedUser, _ := service.UserRepository.FindByEmail(email)
	if existsedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	user := &user.User{
		Name:     name,
		Email:    email,
		Password: "",
	}

	_, err := service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}
