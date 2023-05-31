package dashboard

import (
	"FybBackend/routers/v1/backend/token"
	"FybBackend/routers/v1/exceptionHandler"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
	"time"
)

type MonthData struct {
	Months     string `gorm:"column:months"`
	MonthCount int64  `gorm:"column:monthCount"`
}

func GetPostData(e *gin.Engine, db *gorm.DB) {
	e.GET("/v1/backend/dashboard/getPostData", func(context *gin.Context) {
		if err := token.JwtVerify(context); err != nil {
			context.JSON(403, gin.H{
				"code":    403,
				"message": err.Error(),
			})
			return
		}
		var result *multierror.Error
		var monthData []MonthData
		const countNeed = 6
		endTime := time.Now()
		beginTime := time.Date(time.Now().Year()-1, time.Now().Month(), 1, 0, 0, 0, 0, time.Local)

		err1 := db.Raw("SELECT DATE_FORMAT(publishTime,'%Y-%m') as months , count(*) as monthCount FROM post  where publishTime > ?  and publishTime < ? GROUP BY months", beginTime, endTime).Scan(&monthData).Error

		var data [countNeed]int64
		for i := 0; i < countNeed; i++ {
			tmpTime := time.Date(time.Now().Year(), time.Now().Month()-time.Month(countNeed-i-1), 1, 0, 0, 0, 0, time.Local)
			date := tmpTime.Format("2006-01")
			data[i] = 0
			for j := 0; j < len(monthData); j++ {
				if monthData[j].Months == date {
					data[i] = monthData[j].MonthCount
				}
			}
		}
		result = multierror.Append(result, err1)
		code, msg := exceptionHandler.Handle(result)
		if code == 200 {
			context.JSON(code, gin.H{
				"code":    code,
				"message": "请求成功",
				"data":    data,
			})
		} else {
			context.JSON(code, gin.H{
				"code":    code,
				"message": msg,
			})
		}
	})
}
