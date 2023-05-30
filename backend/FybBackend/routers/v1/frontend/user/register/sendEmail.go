package register

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func SendEmail(e *gin.Engine) {
	e.POST("/v1/frontend/sendEmail", func(context *gin.Context) {
		var result error
		mp := make(map[string]interface{})
		b, err := context.GetRawData()
		if err != nil {
			context.JSON(500, gin.H{
				"code":    500,
				"message": "请求体解析失败",
				"error":   err.Error(),
			})
			return
		}
		if err := json.Unmarshal(b, &mp); err != nil {
			context.JSON(500, gin.H{
				"code":    500,
				"message": "请求体解析失败",
				"error":   err.Error(),
			})
			return
		}

		account := mp["account"].(string)
		vCode := generateVerificationCode()

		c, err := redis.Dial("tcp", ":6379")
		if err != nil {
			context.JSON(500, gin.H{
				"code":    500,
				"message": "无法连接到 Redis",
				"error":   err.Error(),
			})
			return
		}
		defer c.Close()

		_, err = c.Do("set", account, vCode)
		if err != nil {
			result = fmt.Errorf("failed to store verification code in Redis: %w", err)
		}
		_, err = c.Do("expire", account, 300)
		if err != nil {
			result = fmt.Errorf("failed to set expiration for verification code in Redis: %w", err)
		}

		if result != nil {
			context.JSON(500, gin.H{
				"code":    500,
				"message": "验证码发送失败",
				"error":   result.Error(),
			})
			return
		}

		err = SendRegistrationEmail(account, "【福研帮验证码】感谢注册使用福研帮！", generateEmailContent(account, vCode))
		if err != nil {
			context.JSON(500, gin.H{
				"code":    500,
				"message": "验证码发送失败",
				"error":   err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"code":    200,
			"message": "验证码发送成功",
		})
	})
}

func generateVerificationCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%04v", rnd.Int31n(10000))
	return vCode
}

func generateEmailContent(account, vCode string) string {
	t := time.Now().Format("2006-01-02 15:04:05")
	content := fmt.Sprintf(`
尊敬的%s，您好！
			
您于 %s 提交的邮箱验证，本次验证码为%s，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。
			
此邮箱为系统邮箱，请勿回复.
	`, account, t, vCode)
	return content
}
