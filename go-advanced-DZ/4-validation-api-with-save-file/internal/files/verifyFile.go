package files

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileStructure struct {
	Email       string `json:"email" validate:"required,email"`
	Hash_string string `json:"hash" validate:"required"`
}

func SaveVerifyFile(email string, hash_string string) error {
	fmt.Println(email, hash_string)
	newverify := FileStructure{Email: email, Hash_string: hash_string}

	texttofile, err := json.Marshal(newverify)
	if err != nil {
		return err
	}
	nameFile := email + ".json"
	err = os.WriteFile(nameFile, texttofile, 0644)
	if err != nil {
		return err
	}
	return nil
}

func VerifyEmailFile(email string, hash_string string) (bool, error) {
	nameFile := email + ".json"
	data, err := os.ReadFile(nameFile)
	if err != nil {
		return false, err
	}
	err = os.Remove(nameFile)
	if err != nil {
		return false, err
	}
	var fileStruct FileStructure

	err = json.Unmarshal(data, &fileStruct)
	if err != nil {
		return false, err
	}
	if hash_string != fileStruct.Hash_string {
		return true, nil
	}

	return true, nil
}
