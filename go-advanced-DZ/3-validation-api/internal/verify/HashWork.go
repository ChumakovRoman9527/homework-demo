package verify

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Hash_Generate(deps EmailHandler) string {
	// Входные данные
	data := deps.EmailValidationConfig.EmailConfig.Hash_secret
	// Создать новый SHA-256-хэш
	hash := sha256.New()
	// Записать данные в хэш
	data_b := []byte(data)
	hash.Write(data_b)
	// Вычислить хэш
	hashedData := hash.Sum(nil)
	fmt.Println("mySecret:", hex.EncodeToString(hashedData))
	// Преобразовать в шестнадцатеричную строку
	hexHash := hex.EncodeToString(hashedData)
	// Вывести хэш
	return hexHash
}

func Hash_Check(deps EmailHandler, income string) EmailResponse {

	// Вычислить хэш
	hashedData_s := Hash_Generate(deps)
	hashedData := []byte(hashedData_s)

	income_b := []byte(income)
	fmt.Println("myhash :", hashedData_s)
	fmt.Println("outhash:", income)

	fmt.Println("mySecret 0:", hex.EncodeToString(hashedData))
	fmt.Println("mySecret 1:", hex.EncodeToString(income_b))
	if bytes.Compare(hashedData, income_b) != 0 {

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
