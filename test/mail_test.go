package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"light-cloud/src/core/define"
	"net/smtp"
	"testing"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Light Cloud <发送邮箱>"
	e.To = []string{"接受邮箱"}
	e.Subject = "发送邮件测试"
	e.HTML = []byte("你的验证码为：<h1>1234567</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", define.MailUsername, define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"},
	)
	if err != nil {
		t.Fatal(err)
	}
}
