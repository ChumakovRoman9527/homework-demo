package auth

import (
	"6-order-api-cart/internal/user"
	"errors"
	"math/rand"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	PhoneAuthRepository *PhoneAuthRepository
	UserRepository      *user.UserRepository
}

func NewAuthService(authRepository *PhoneAuthRepository, userRepository *user.UserRepository) *AuthService {
	return &AuthService{PhoneAuthRepository: authRepository, UserRepository: userRepository}
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

	/*Вот здесь попытка создать пользователя по номеру телефона*/
	user_id, err := service.UserRepository.GetIdByPhone(phone)
	if err != nil {
		return "", err
	}
	/*Вот здесь попытка создать пользователя по номеру телефона*/

	newAuthfromReq := &PhoneAuth{
		SessionID: string(sessionId),
		Code:      code,
		Phone:     phone,
		UserId:    uint(user_id),
	}

	newAuth, err := service.PhoneAuthRepository.Create(newAuthfromReq)
	if err != nil {
		return "", err
	}
	return newAuth.SessionID, nil

}

func (service AuthService) CodeCheck(sessionId string, code string) (string, error) {

	existed, err := service.PhoneAuthRepository.GetBySessionCode(sessionId, code)
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
