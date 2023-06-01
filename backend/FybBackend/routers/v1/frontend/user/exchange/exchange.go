package exchange

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"github.com/jordan-wright/email"
	"net/http"
	"net/smtp"
)

func Exchange(e *gin.Engine) {
	db := fybDatabase.InitDB()
	e.POST("/v1/frontend/user/exchange", func(context *gin.Context) {
		var result *multierror.Error

		data, err := context.GetRawData()
		if err != nil {
			result = multierror.Append(result, err)
		}
		var m map[string]interface{}
		err = json.Unmarshal(data, &m)
		if err != nil {
			result = multierror.Append(result, err)
		}

		account, flag, err := fybDatabase.DecreaseBalance(db, int64(m["userID"].(float64)), int64(m["amount"].(float64)))

		if err != nil {
			result = multierror.Append(result, err)
		}

		if result.ErrorOrNil() == nil && flag == true {
			subject := fmt.Sprintf("%s 兑换了 %s", account, m["name"].(string))
			content := fmt.Sprintf("%s 兑换了 %s", account, m["name"].(string))
			sendEmail("dhzemail@foxmail.com", subject, content)
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "登记成功",
			})
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": result.Error(),
			})
		}
	})
}

func sendEmail(to, subject, body string) error {
	e := email.NewEmail()
	e.From = "福研帮_Official <dhzemail@foxmail.com>"
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(body)

	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "dhzemail@foxmail.com", "mzvarxlhnhjydida", "smtp.qq.com"))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
