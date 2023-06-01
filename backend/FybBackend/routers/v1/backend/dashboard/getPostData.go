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

type HourData struct {
	Hours     int   `gorm:"column:hours"`
	HourCount int64 `gorm:"column:hourCount"`
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
		var hourData []HourData

		const monthNeed = 6
		const periodNeed = 6

		endTime := time.Now()
		beginTime := time.Date(time.Now().Year()-1, time.Now().Month(), 1, 0, 0, 0, 0, time.Local)

		err1 := db.Raw("SELECT DATE_FORMAT(publishTime,'%Y-%m') as months , count(*) as monthCount FROM post  where publishTime > ?  and publishTime < ? GROUP BY months", beginTime, endTime).Scan(&monthData).Error
		var monthPostNum [monthNeed]int64
		for i := 0; i < monthNeed; i++ {
			tmpTime := time.Date(time.Now().Year(), time.Now().Month()-time.Month(monthNeed-i-1), 1, 0, 0, 0, 0, time.Local)
			date := tmpTime.Format("2006-01")
			monthPostNum[i] = 0
			for j := 0; j < len(monthData); j++ {
				if monthData[j].Months == date {
					monthPostNum[i] = monthData[j].MonthCount
				}
			}
		}
		endTime = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
		beginTime = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-1, 0, 0, 0, 0, time.Local)

		err2 := db.Raw("SELECT DATE_FORMAT(publishTime,'%H') as hours , count(*) as  hourCount FROM post  where publishTime > ?  and publishTime < ? GROUP BY hours", beginTime, endTime).Scan(&hourData).Error
		var hourPostNum [periodNeed]int64
		for i := 0; i < periodNeed; i++ {
			beginHour := i * (24 / periodNeed)
			endHour := (i + 1) * (24 / periodNeed)
			hourPostNum[i] = 0
			for j := 0; j < len(hourData); j++ {
				if hourData[j].Hours > beginHour && hourData[j].Hours < endHour {
					hourPostNum[i]++
				}
			}
		}

		result = multierror.Append(result, err1, err2)
		code, msg := exceptionHandler.Handle(result)
		if code == 200 {
			context.JSON(code, gin.H{
				"code":    code,
				"message": "请求成功",
				"data": map[string]interface{}{
					"monthPostNum": monthPostNum,
					"dayPostNum":   hourPostNum,
				},
			})
		} else {
			context.JSON(code, gin.H{
				"code":    code,
				"message": msg,
			})
		}
	})
}
