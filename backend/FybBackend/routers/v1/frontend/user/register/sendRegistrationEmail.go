package register

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendRegistrationEmail(account, subject string, content string) error {
	e := email.NewEmail()
	e.From = "福研帮_Official <dhzemail@foxmail.com>"
	e.To = []string{account}
	e.Subject = subject
	e.Text = []byte(content)

	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "dhzemail@foxmail.com", "mzvarxlhnhjydida", "smtp.qq.com"))
	if err != nil {
		// 处理邮件发送失败的错误
		return fmt.Errorf("failed to send registration email: %w", err)
	}

	return nil
}
