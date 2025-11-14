package auth

import (
	"11-JWTAUTH/internal/user"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := &user.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}
	fmt.Println(email, password, name, user)
	_, err = service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}
