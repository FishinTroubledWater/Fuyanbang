package addUser

import (
	fybDatabase "FybBackend/database"
	"FybBackend/routers/v1/backend/user/token"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

type responseItem struct {
	Id      int64  `json:"id"`
	Account string `json:"account"`
	Token   string `json:"token"`
}

func AddUser(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/backend/user/add", func(context *gin.Context) {
		var errors *multierror.Error
		mp := make(map[string]interface{})
		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)
		err3 := token.JwtVerify(context)
		err4 := fybDatabase.AddUser(db, mp)
		multierror.Append(errors, err1, err2, err3, err4)
		if errors.ErrorOrNil() == nil {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "add user success!",
			})
		} else {
			context.JSON(404, gin.H{
				"code":    404,
				"message": errors.Error(),
			})
		}
	})
}
