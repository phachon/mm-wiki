package utils

import (
	"crypto/tls"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/phachon/mm-wiki/global"
	"gopkg.in/gomail.v2"
	"gopkg.in/russross/blackfriday.v2"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

var Email = NewEmail()

type email struct {
}

func NewEmail() *email {
	return &email{}
}

func (e *email) Send(emailConf map[string]string, toList []string, subject string, body string) error {

	from := emailConf["sender_address"]
	if emailConf["sender_name"] != "" {
		from = fmt.Sprintf("%s <%s>", emailConf["sender_name"], emailConf["sender_address"])
	}
	var tt []*gomail.Message
	for _, toAddress := range toList {
		m := gomail.NewMessage()
		m.SetHeader("From", from)
		m.SetHeader("To", toAddress)
		m.SetHeader("Subject", emailConf["sender_title_prefix"]+subject)
		m.SetBody("text/html", body)
		tt = append(tt, m)
	}
	portInt, _ := strconv.Atoi(emailConf["port"])
	d := gomail.NewDialer(emailConf["host"], portInt, emailConf["username"], emailConf["password"])
	if emailConf["is_ssl"] == "1" {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
	return d.DialAndSend(tt...)
}

func (e *email) SendByEmail(email map[string]string, to []string, subject string, body string, contentType string) error {

	userEmail := email["sender_address"]
	addr := fmt.Sprintf("%s:%s", email["host"], email["port"])
	auth := smtp.PlainAuth("", userEmail, email["password"], email["host"])
	user := email["username"]
	nickname := email["sender_name"]
	subject = email["sender_title_prefix"] + subject

	msg := fmt.Sprintf("To: %s \r\nFrom: %s <%s>\r\nSubject: %s \r\nContent-Type: text/%s; charset=UTF-8\r\n\r\n%s",
		strings.Join(to, ","), nickname, user, subject, contentType, body)

	err := smtp.SendMail(addr, auth, user, to, []byte(msg))
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
	body = strings.Replace(body, "{{.copyright}}", global.SYSTEM_COPYRIGHT, 1)

	content := string(blackfriday.Run([]byte(document["content"])))
	body = strings.Replace(body, "{{.document_content}}", content, 1)

	return
}
