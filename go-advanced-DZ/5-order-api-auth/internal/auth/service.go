package auth

import "math/rand"

type AuthService struct {
	UserRepository *PhoneAuthRepository
}

func NewAuthService(authRepository *PhoneAuthRepository) *AuthService {
	return &AuthService{UserRepository: authRepository}
}

const codeLen = 4

func newCode() string {
	var res string
	var rndStr string
	for i := 1; i <= codeLen; i++ {
		rndStr = string(rand.Intn(10))
		res += rndStr
	}
	return res
}

func (service AuthService) CodeGenerate(phone string) (string, error) {
	code := newCode()

}
