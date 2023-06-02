package register

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-multierror"
	"time"
)

func ValidateEmailCode(e *gin.Engine) {
	e.POST("/v1/frontend/codeVerify", func(context *gin.Context) {
		var result error
		flag := false

		mp := make(map[string]interface{})
		b, err := context.GetRawData()
		if err != nil {
			result = multierror.Append(result, fmt.Errorf("failed to read request body: %w", err))
		}
		if err := json.Unmarshal(b, &mp); err != nil {
			result = multierror.Append(result, fmt.Errorf("failed to unmarshal JSON: %w", err))
		}

		vCode := mp["code"]
		account := mp["account"]

		c, err := redis.Dial("tcp", ":6379", redis.DialConnectTimeout(time.Second))
		if err != nil {
			context.JSON(500, gin.H{
				"code":    500,
				"message": "无法连接到 Redis",
				"error":   err.Error(),
			})
			return
		}
		defer c.Close()

		vCodeRaw, err := redis.Bytes(c.Do("GET", account))
		if err != nil {
			if err != redis.ErrNil {
				result = multierror.Append(result, fmt.Errorf("failed to retrieve verification code from Redis: %w", err))
			}
		} else {
			if vCodeRaw != nil && string(vCodeRaw) != "" && vCode == string(vCodeRaw) {
				flag = true
			}
		}

		if result == nil && flag {
			// 发送注册成功邮件
			content := fmt.Sprintf(`
亲爱的%s，您好！
			
在你踏上考研之路时，福研帮_Official愿与你并肩同行，为你提供全方位的支持和陪伴。无论你面临怎样的困难和挑战，记住你不是孤单一人，我们将一直在你身旁。

在这艰辛的备考过程中，努力坚持，不放弃。相信自己的能力和潜力，勇往直前。无论遇到怎样的瓶颈和困扰，不要气馁，因为你有福研帮作为你的后盾。

福研帮_Official衷心祝愿你能够充满自信和毅力，攀登考研的高峰。愿你在考试中发挥出自己最好的水平，取得令人骄傲的成绩。无论结果如何，记住努力的过程才是最宝贵的财富。

加油，同学！福研帮_Official将一直与你同行，为你的考研之路加油助力。愿你实现理想，迈向美好未来！

祝福考研上岸！

福研帮_Official
	`, account)
			err := SendRegistrationEmail(account.(string), "【福研帮】系统消息，请勿回复", content)
			if err != nil {
				// 处理邮件发送失败的错误
				context.JSON(500, gin.H{
					"code":    500,
					"message": "邮件发送失败",
					"error":   err.Error(),
				})
			}

			context.JSON(200, gin.H{
				"code":    200,
				"message": "验证成功",
			})

		} else {
			errMsg := ""
			if result != nil {
				errMsg = result.Error()
			}
			context.JSON(404, gin.H{
				"code":    404,
				"message": "验证码错误或已失效",
				"error":   errMsg,
			})
		}
	})
}
