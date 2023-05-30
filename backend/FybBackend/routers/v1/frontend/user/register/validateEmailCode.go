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
			content := fmt.Sprintf("欢迎加入福研帮！%s", account.(string))
			err := SendRegistrationEmail(account.(string), "【福研帮】注册成功！欢迎加入福研帮！", content)
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
