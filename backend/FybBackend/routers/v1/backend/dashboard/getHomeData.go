package dashboard

import (
	fybDatabase "FybBackend/database"
	"FybBackend/routers/v1/backend/token"
	"FybBackend/routers/v1/exceptionHandler"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
	"time"
)

type HomeData struct {
	PhoneNumber     string `json:"phoneNumber"`
	NumberOfNewUser int64  `json:"numberOfNewUser"`
	NumberOfNewPost int64  `json:"numberOfNewPost"`
	Reviewed        int64  `json:"reviewed"`
	UserNumber      int64  `json:"userNumber"`
	PostNumber      int64  `json:"postNumber"`
	FeedbackNumber  int64  `json:"feedbackNumber"`
}

func GetHomeData(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/backend/dashboard/getHomeData", func(context *gin.Context) {
		if err := token.JwtVerify(context); err != nil {
			context.JSON(403, gin.H{
				"code":    403,
				"message": err.Error(),
			})
			return
		}
		mp := make(map[string]interface{})
		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)
		admin, _, err3 := fybDatabase.SelectSingleAdminByCondition(db, mp)
		var result *multierror.Error

		var numberOfNewUser int64 = 0
		var numberOfNewPost int64 = 0
		var reviewed int64 = 0
		var userNumber int64 = 0
		var postNumber int64 = 0
		var feedbackNumber int64 = 0
		beginTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-1, 0, 0, 0, 0, time.Local)
		endTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
		err4 := db.Raw("SELECT count(*) as numberOfNewUser from user where registerTime > ? and registerTime < ?", beginTime, endTime).Scan(&numberOfNewUser).Error
		err5 := db.Raw("SELECT count(*) as numberOfNewPost from post where publishTime > ? and publishTime < ?", beginTime, endTime).Scan(&numberOfNewPost).Error
		err6 := db.Raw("SELECT count(*) as numberOfNewUser from post where state = 0").Scan(&reviewed).Error
		err7 := db.Raw("SELECT count(*) as userNumber from user").Scan(&userNumber).Error
		err8 := db.Raw("SELECT count(*) as postNumber from post").Scan(&postNumber).Error
		err9 := db.Raw("SELECT count(*) as feedbackNumber from feedback").Scan(&feedbackNumber).Error
		result = multierror.Append(result, err1, err2, err3, err4, err5, err6, err7, err8, err9)
		code, msg := exceptionHandler.Handle(result)
		if code == 200 {
			context.JSON(code, gin.H{
				"code":    code,
				"message": "请求成功",
				"data": HomeData{
					PhoneNumber:     admin.PhoneNumber,
					NumberOfNewUser: numberOfNewUser,
					NumberOfNewPost: numberOfNewPost,
					Reviewed:        reviewed,
					UserNumber:      userNumber,
					PostNumber:      postNumber,
					FeedbackNumber:  feedbackNumber,
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
