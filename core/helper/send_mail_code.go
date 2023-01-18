package helper

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"light-cloud/src/core/define"
	"math/rand"
	"net/smtp"
	"strings"
	"time"
)

// SendMailCode 发送邮箱验证码
func SendMailCode(mail string) error {
	e := email.NewEmail()
	e.From = "Light Cloud <" + define.MailUsername + ">"
	e.To = []string{mail}
	e.Subject = "👏欢迎使用Light Cloud"
	e.HTML = []byte("您的验证码为：<h1>" + GenValidateCode() + ", 请确保是本人操作，请勿泄漏您的验证码</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", define.MailUsername, define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"},
	)
	if err != nil {
		return err
	}

	return nil
}

// GenValidateCode 随机生成8位验证码
func GenValidateCode() string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < 8; i++ {
		_, err := fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
		if err != nil {
			return ""
		}
	}
	return sb.String()
}
