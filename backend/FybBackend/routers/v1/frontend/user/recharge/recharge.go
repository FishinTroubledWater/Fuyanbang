package recharge

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"net/http"
)

func Recharge(e *gin.Engine) {
	db := fybDatabase.InitDB()
	e.POST("/v1/frontend/user/recharge", func(context *gin.Context) {
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

		flag, err := fybDatabase.AddBalance(db, int64(m["userID"].(float64)), int64(m["amount"].(float64)))

		if err != nil {
			result = multierror.Append(result, err)
		}

		if result.ErrorOrNil() == nil && flag == true {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "充值成功",
			})
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": result.Error(),
			})
		}
	})
}
