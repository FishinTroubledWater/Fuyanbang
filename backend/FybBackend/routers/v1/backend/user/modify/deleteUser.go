package modify

import (
	fybDatabase "FybBackend/database"
	"FybBackend/routers/v1/backend/user/token"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func DeleteUser(e *gin.Engine, db *gorm.DB) {
	e.DELETE("/v1/backend/user/delete", func(context *gin.Context) {
		var result *multierror.Error
		mp := make(map[string]interface{})
		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)
		delete(mp, "phoneNumber")
		err3 := token.JwtVerify(context)
		_, err4 := fybDatabase.DeleteUser(db, mp)
		result = multierror.Append(result, err1, err2, err3, err4)
		if result.ErrorOrNil() == nil {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "删除成功",
			})
		} else {
			var code int
			if err3 != nil {
				code = 403
			} else if err4 != nil {
				code = 404
			} else {
				code = 500
			}
			context.JSON(code, gin.H{
				"code":    code,
				"message": result.Error(),
			})

		}
	})
}
