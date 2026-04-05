package mail

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"strings"
)

// MailConfig 邮件发送配置
type MailConfig struct {
	Host              string // SMTP 服务器主机
	Port              int    // SMTP 服务器端口
	Username          string // 用户名
	Password          string // 密码
	SenderAddress     string // 发件人地址
	SenderName        string // 发件人显示名
	SenderTitlePrefix string // 邮件标题前缀
	IsSSL             bool   // 是否使用SSL/TLS
}

// MailMessage 邮件内容
type MailMessage struct {
	To      []string // 收件人列表
	Subject string   // 邮件主题
	Body    string   // 邮件正文
	IsHTML  bool     // 是否 HTML 格式
}

// SendMail 发送邮件
func SendMail(config *MailConfig, message *MailMessage) error {
	if config == nil {
		return fmt.Errorf("mail config is nil")
	}
	if message == nil {
		return fmt.Errorf("mail message is nil")
	}
	if len(message.To) == 0 {
		return fmt.Errorf("mail to is empty")
	}

	// 构建邮件内容
	subject := message.Subject
	if config.SenderTitlePrefix != "" {
		subject = config.SenderTitlePrefix + subject
	}

	contentType := "text/plain"
	if message.IsHTML {
		contentType = "text/html"
	}

	from := config.SenderAddress
	if config.SenderName != "" {
		from = fmt.Sprintf("%s <%s>", config.SenderName, config.SenderAddress)
	}

	header := make(map[string]string)
	header["From"] = from
	header["To"] = strings.Join(message.To, ",")
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=UTF-8", contentType)

	body := ""
	for k, v := range header {
		body += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	body += "\r\n" + message.Body

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)

	if config.IsSSL {
		return sendMailWithTLS(addr, auth, config, message.To, []byte(body))
	}
	return smtp.SendMail(addr, auth, config.SenderAddress, message.To, []byte(body))
}

// sendMailWithTLS 使用 TLS 发送邮件
func sendMailWithTLS(addr string, auth smtp.Auth, config *MailConfig, to []string, msg []byte) error {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return fmt.Errorf("split host port err: %s", err.Error())
	}

	tlsConfig := &tls.Config{
		ServerName: host,
	}

	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return fmt.Errorf("tls dial err: %s", err.Error())
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("new smtp client err: %s", err.Error())
	}
	defer client.Close()

	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("smtp auth err: %s", err.Error())
	}

	if err = client.Mail(config.SenderAddress); err != nil {
		return fmt.Errorf("smtp mail err: %s", err.Error())
	}

	for _, addr := range to {
		if err = client.Rcpt(addr); err != nil {
			return fmt.Errorf("smtp rcpt err: %s", err.Error())
		}
	}

	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("smtp data err: %s", err.Error())
	}

	_, err = w.Write(msg)
	if err != nil {
		return fmt.Errorf("smtp write err: %s", err.Error())
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("smtp close err: %s", err.Error())
	}

	return client.Quit()
}

// TestConnection 测试邮件服务器连接
func TestConnection(config *MailConfig) error {
	if config == nil {
		return fmt.Errorf("mail config is nil")
	}

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	if config.IsSSL {
		host, _, err := net.SplitHostPort(addr)
		if err != nil {
			return fmt.Errorf("split host port err: %s", err.Error())
		}

		tlsConfig := &tls.Config{
			ServerName: host,
		}

		conn, err := tls.Dial("tcp", addr, tlsConfig)
		if err != nil {
			return fmt.Errorf("tls dial err: %s", err.Error())
		}
		defer conn.Close()

		client, err := smtp.NewClient(conn, host)
		if err != nil {
			return fmt.Errorf("new smtp client err: %s", err.Error())
		}
		defer client.Close()

		auth := smtp.PlainAuth("", config.Username, config.Password, host)
		if err = client.Auth(auth); err != nil {
			return fmt.Errorf("smtp auth err: %s", err.Error())
		}
		return client.Quit()
	}

	client, err := smtp.Dial(addr)
	if err != nil {
		return fmt.Errorf("smtp dial err: %s", err.Error())
	}
	defer client.Close()

	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("smtp auth err: %s", err.Error())
	}
	return client.Quit()
}
