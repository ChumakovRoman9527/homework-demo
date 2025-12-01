package auth

import (
	"errors"
	"math/rand"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

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
		rndStr = strconv.Itoa(rand.Intn(10))
		res += rndStr
	}
	return res
}

func (service AuthService) SessionGenerate(phone string) (string, error) {

	code := newCode()

	//Тут поидее мы должны уже обратится к сервису отправки sms

	forsession := phone + "|" + code

	sessionId, err := bcrypt.GenerateFromPassword([]byte(forsession), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	newAuthfromReq := &PhoneAuth{
		SessionID: string(sessionId),
		Code:      code,
		Phone:     phone,
	}

	newAuth, err := service.UserRepository.Create(newAuthfromReq)
	if err != nil {
		return "", err
	}
	return newAuth.SessionID, nil

}

func (service AuthService) CodeCheck(sessionId string, code string) (string, error) {

	existed, err := service.UserRepository.GetBySessionCode(sessionId, code)
	if err != nil {
		return "", err
	}
	if existed == nil {
		return "", errors.New(ErrUserWrongCredentials)
	}
	return existed.Phone, nil
}

func (service AuthService) ClearOldSession(sessionId string) error {

	// _, err := service.UserRepository.DeleteBySessionCode(sessionId)
	// if err != nil {
	// 	return err
	// }

	return nil
}
