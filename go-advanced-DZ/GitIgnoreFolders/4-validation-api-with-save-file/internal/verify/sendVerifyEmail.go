package verify

import (
	"4-validation-api-with-save-file/internal/files"
	"fmt"
	"net/smtp"
	"time"

	//модуль нормально не импортировался, пришлось алиасить
	another_email "github.com/jordan-wright/email"
)

type EmailRecipient struct {
	RecipientEmail string `json:"RecipientEmail" validate:"required,email"`
	RecipientName  string `json:"RecipientName" validate:"required"`
}

func SendVerifyEmail(deps EmailHandler, body *EmailRegister) EmailResponse {
	emailFrom := deps.EmailConfig.Email
	pass := deps.EmailConfig.Password
	smtpServ := deps.EmailConfig.Address
	port := deps.EmailConfig.Port
	fmt.Println("r.Body:", body)

	// fmt.Println(t.RecipientEmail)
	// fmt.Println(t.RecipientName)

	smtpServPort := smtpServ + ":" + port

	hashString := Hash_Generate(body.Email, deps, "")

	_, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		fmt.Println("Не удачная попытка установки временной зоны ! ->", err.Error())

	}
	// fmt.Println(time.Now().Local().String())

	emailString := "<a href=http://localhost:8081/verify/" + hashString + ">Verify !</a>"
	e := another_email.NewEmail()
	e.From = "localhost@localhost.local"
	e.To = []string{body.Email}
	e.Bcc = []string{body.Email}
	e.Cc = []string{body.Email}
	e.Subject = "VerifyEmail for " + body.Name
	e.HTML = []byte("<div><h1>Verify the email !!</h1><p> " + emailString + "</p><p>" + time.Now().Local().String() + "</p></div>")
	err = e.Send(smtpServPort, smtp.PlainAuth("", emailFrom, pass, smtpServ))
	if err != nil {
		return EmailResponse{
			statusCode: 500,
			statusText: err.Error(),
		}
	}

	err = files.SaveVerifyFile(body.Email, hashString)
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
