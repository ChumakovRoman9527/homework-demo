package verify

import (
	"4-validation-api-with-save-file/internal/files"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Hash_Generate(email string, deps EmailHandler, _dt string) string {
	_, _ = time.LoadLocation("Europe/Moscow")

	dt := time.Now().Unix()
	if _dt != "" {
		dt, _ = strconv.ParseInt(_dt, 10, 64)

	}
	// fmt.Println("_dt = ", _dt)
	// fmt.Println("dt = ", dt)

	// Входные данные
	data := fmt.Sprintf("%s|||%d|||%s", email, dt, deps.EmailValidationConfig.EmailConfig.Hash_secret)
	// Создать новый SHA-256-хэш
	hash := sha256.Sum256([]byte(data))

	hexHash := base64.URLEncoding.EncodeToString([]byte(email + ":" + fmt.Sprintf("%d", dt) + ":" + hex.EncodeToString(hash[:])))
	// Вывести хэш
	return hexHash
}

func Hash_Check(deps EmailHandler, income string) EmailResponse {

	decoded, err := base64.URLEncoding.DecodeString(income)
	if err != nil {
		return EmailResponse{
			statusCode: 500,
			statusText: err.Error(),
		}
	}
	parts := strings.Split(string(decoded), ":")
	if len(parts) != 3 {
		return EmailResponse{statusCode: 400, statusText: "Неверный формат hash"}
	}
	email := parts[0]
	// dt := parts[1]
	// receivedHash := parts[2]

	success, err := files.VerifyEmailFile(email, income)
	if err != nil {
		return EmailResponse{
			statusCode: 500,
			statusText: err.Error(),
		}
	}

	// expectedHash_URL := Hash_Generate(email, deps, dt)

	// expecteddecoded, _ := base64.URLEncoding.DecodeString(expectedHash_URL)
	// expectedparts := strings.Split(string(expecteddecoded), ":")
	// expectedHash := expectedparts[2]

	if !success {

		return EmailResponse{
			statusCode: 500,
			statusText: "Hash неверен !",
		}
	}
	return EmailResponse{
		statusCode: 200,
		statusText: "Hash совпал",
	}
}
