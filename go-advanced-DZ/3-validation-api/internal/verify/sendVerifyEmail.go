package verify

import (
	"net/smtp"

	//модуль нормально не импортировался, пришлось алиасить
	another_email "github.com/jordan-wright/email"
)

func SendVerifyEmail(deps EmailHandler) EmailResponse {
	emailFrom := deps.EmailConfig.Email
	pass := deps.EmailConfig.Password
	smtpServ := deps.EmailConfig.Address
	port := deps.EmailConfig.Port

	smtpServPort := smtpServ + ":" + port

	hashString := Hash_Generate(deps)

	emailString := "<a href=http://localhost:8081/verify/" + hashString + ">Verify !</a>"
	e := another_email.NewEmail()
	e.From = "rvch <rvch84@gmail.com>"
	e.To = []string{"rvch84@gmail.com"}
	e.Bcc = []string{"rvch84@gmail.com"}
	e.Cc = []string{"rvch84@gmail.com"}
	e.Subject = "VerifyEmail"
	e.HTML = []byte("<h1>Verify the email !!" + emailString + "</h1>")
	err := e.Send(smtpServPort, smtp.PlainAuth("", emailFrom, pass, smtpServ))
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
