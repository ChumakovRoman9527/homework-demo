package verify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"time"

	//модуль нормально не импортировался, пришлось алиасить
	another_email "github.com/jordan-wright/email"
)

type EmailRecipient struct {
	RecipientEmail string `json:"RecipientEmail"`
	RecipientName  string `json:"RecipientName"`
}

func SendVerifyEmail(deps EmailHandler, r *http.Request) EmailResponse {
	emailFrom := deps.EmailConfig.Email
	pass := deps.EmailConfig.Password
	smtpServ := deps.EmailConfig.Address
	port := deps.EmailConfig.Port

	decoder := json.NewDecoder(r.Body)
	var t EmailRecipient
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Println("Не удачная попытка разбора JSON ! ->", err.Error())
		return EmailResponse{
			statusCode: 500,
			statusText: err.Error(),
		}
	}
	// fmt.Println(t.RecipientEmail)
	// fmt.Println(t.RecipientName)

	smtpServPort := smtpServ + ":" + port

	hashString := Hash_Generate(deps)

	_, err = time.LoadLocation("Europe/Moscow")
	if err != nil {
		fmt.Println("Не удачная попытка установки временной зоны ! ->", err.Error())

	}
	fmt.Println(time.Now().Local().String())

	emailString := "<a href=http://localhost:8081/verify/" + hashString + ">Verify !</a>"
	e := another_email.NewEmail()
	e.From = "localhost@localhost.local"
	e.To = []string{t.RecipientEmail}
	e.Bcc = []string{t.RecipientEmail}
	e.Cc = []string{t.RecipientEmail}
	e.Subject = "VerifyEmail for " + t.RecipientName
	e.HTML = []byte("<div><h1>Verify the email !!</h1><p> " + emailString + "</p><p>" + time.Now().Local().String() + "</p></div>")
	err = e.Send(smtpServPort, smtp.PlainAuth("", emailFrom, pass, smtpServ))
	if err != nil {
		return EmailResponse{
			statusCode: 500,
			statusText: err.Error(),
		}
	}
	return EmailResponse{
		statusCode: 200,
		statusText: "Email отправлен",
	}
}
