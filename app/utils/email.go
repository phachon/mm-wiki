package utils

import (
	"strings"
	"net/smtp"
	"time"
	"gopkg.in/russross/blackfriday.v2"
	"github.com/astaxie/beego"
)

var Email = NewEmail()

type email struct {

}

func NewEmail() *email {
	return &email{}
}

func (e *email) SendByEmail(email map[string]string, to []string, subject string, body string, contentType string) error {

	userEmail := email["sender_address"]
	smtpPort := ":"+email["port"]
	mailPassword := email["password"]
	smtpHost := email["host"]
	auth := smtp.PlainAuth("", userEmail, mailPassword, smtpHost)
	nickname := email["sender_name"]
	user := email["username"]

	subject = email["sender_title_prefix"]+subject
	contentType = "Content-Type: text/"+contentType+"; charset=UTF-8"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	err := smtp.SendMail(smtpHost+smtpPort, auth, user, to, msg)
	return err
}

func (e *email) MakeDocumentHtmlBody(document map[string]string, view string) (body string, err error) {
	viewTemplate, err := File.GetFileContents(view)
	if err != nil {
		return
	}

	body = strings.Replace(viewTemplate, "{{.now_time}}", beego.Date(time.Now(), "Y-m-d H:i:s"), 1)
	body = strings.Replace(body, "{{.document_name}}", document["name"], 1)
	body = strings.Replace(body, "{{.username}}", document["username"], 1)
	body = strings.Replace(body, "{{.update_time}}", Date.Format(document["update_time"], "Y-m-d H:i:s"), 1)
	body = strings.Replace(body, "{{.comment}}", document["comment"], 1)
	body = strings.Replace(body, "{{.document_url}}", document["url"], 1)

	content := string(blackfriday.Run([]byte(document["content"])))
	body = strings.Replace(body, "{{.document_content}}", content, 1)

	return
}