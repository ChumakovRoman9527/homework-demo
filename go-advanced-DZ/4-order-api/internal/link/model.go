package link

import (
	"math/rand"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"UniqueIndex"`
}

type Product struct {
	gorm.Model
	Name        string         `json:"Name"`
	Description string         `json:"Description"`
	Images      pq.StringArray `json:"Image" gorm:"type:text[]"`
}

func NewLink(url string) *Link {
	return &Link{
		Url:  url,
		Hash: RabdStringRunes(6),
	}
}

func getEnglishAlphabet() []rune {
	var alphabet []rune

	// Добавляем заглавные буквы A-Z (коды Unicode: 65–90)
	for i := 65; i <= 90; i++ {
		alphabet = append(alphabet, rune(i))
	}

	// Добавляем строчные буквы a-z (коды Unicode: 97–122)
	for i := 97; i <= 122; i++ {
		alphabet = append(alphabet, rune(i))
	}

	return alphabet
}
func RabdStringRunes(n int) string {
	b := make([]rune, n)
	letterRunes := getEnglishAlphabet()
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
