package circle

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func PostAnswerStatus(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/frontend/circle/postAnswerStatus", func(context *gin.Context) {
		var result *multierror.Error
		mp := make(map[string]interface{})

		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)

		queId, err := strconv.ParseInt(mp["queId"].(string), 10, 64)
		if err != nil {
			result = multierror.Append(result, err)
		}

		answerId, err := strconv.ParseInt(mp["answerId"].(string), 10, 64)
		if err != nil {
			result = multierror.Append(result, err)
		}

		flag, err3 := fybDatabase.AddAdoptRecord(db, queId, answerId)

		result = multierror.Append(result, err1, err2, err3)
		if result.ErrorOrNil() == nil && flag == true {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "请求成功",
			})
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": result.Error(),
			})
		}
	})
}
