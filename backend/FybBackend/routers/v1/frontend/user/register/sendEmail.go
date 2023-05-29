package register

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"github.com/jordan-wright/email"
	"log"
	"math/rand"
	"net/smtp"
	"time"
)

func SendEmail(e *gin.Engine) {
	e.POST("/v1/frontend/sendEmail", func(context *gin.Context) {

		var result *multierror.Error
		mp := make(map[string]interface{})
		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)
		result = multierror.Append(result, err1, err2)
		var em []string
		em = append(em, mp["account"].(string))

		e := email.NewEmail()
		e.From = fmt.Sprintf("福研帮_Official <dhzemail@foxmail.com>")
		e.To = em

		// 生成4位随机验证码
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		vCode := fmt.Sprintf("%04v", rnd.Int31n(10000))
		t := time.Now().Format("2006-01-02 15:04:05")

		//设置文件发送的内容
		content := fmt.Sprintf(`
			
尊敬的%s，您好！
			
您于 %s 提交的邮箱验证，本次验证码为%s，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。
			
此邮箱为系统邮箱，请勿回复。

	`, em[0], t, vCode)

		e.Text = []byte(content)
		//设置服务器相关的配置
		err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "dhzemail@foxmail.com", "mzvarxlhnhjydida", "smtp.qq.com"))
		result = multierror.Append(result, err)

		c, err := redis.Dial("tcp", ":6379")
		if err != nil {
			fmt.Printf("redis.Dial() error:%v", err)
			result = multierror.Append(result, err)
		}
		defer c.Close()

		_, err = c.Do("set", "vCode", vCode)
		_, _ = c.Do("expire", "vCode", 300)
		if err != nil {
			log.Println(err.Error())
			result = multierror.Append(result, err)
		}

		if result.ErrorOrNil() == nil {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "验证码发送成功",
			})
		} else {
			context.JSON(404, gin.H{
				"message": 404,
				"msg":     result.Error(),
			})
		}
	})
}
