package dashboard

import (
	"FybBackend/routers/v1/backend/token"
	"FybBackend/routers/v1/exceptionHandler"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
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
		var data []MonthData
		var err1 error
		err1 = db.Raw("SELECT DATE_FORMAT(publishTime,'%Y%m') as months , count(*) as monthCount FROM post GROUP BY months").Scan(&data).Error
		//for rows.Next() {
		//	var monthData MonthData
		//	rows.Scan(&monthData.Months, &monthData.MonthCount)
		//	fmt.Println(monthData)
		//}
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
