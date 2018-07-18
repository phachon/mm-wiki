package utils

import (
	"strings"
	"net/smtp"
)

var Email = NewEmail()

type email struct {

}

func NewEmail() *email {
	return &email{}
}

func (e *email) SendByEmail(email map[string]string, to []string, subject string, body string) error {

	userEmail := email["sender_address"]
	smtpPort := ":"+email["port"]
	mailPassword := email["password"]
	smtpHost := email["host"]
	auth := smtp.PlainAuth("", userEmail, mailPassword, smtpHost)
	//to := []string{"295009256@qq.com"}
	nickname := email["sender_name"]
	user := email["username"]

	subject = email["sender_title_prefix"]+subject
	contentType := "Content-Type: text/plain; charset=UTF-8"
	//body := "邮件内容."
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	err := smtp.SendMail(smtpHost+smtpPort, auth, user, to, msg)
	return err
}

func (e *email) MakeMsg(to []string, nickname string, user string, subject string, body string) []byte {

	contentType := "Content-Type: text/plain; charset=UTF-8"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)

	return msg
}