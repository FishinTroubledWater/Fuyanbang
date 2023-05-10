package modifyUser

import (
	fybDatabase "FybBackend/database"
	"FybBackend/routers/v1/backend/token"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func UpdateUser(e *gin.Engine, db *gorm.DB) {
	e.PATCH("/v1/backend/user/update", func(context *gin.Context) {
		var result *multierror.Error
		mp1 := make(map[string]interface{})
		mp2 := make(map[string]interface{})
		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp1)
		mp2["phoneNumber"] = mp1["phoneNumber"]
		delete(mp1, "phoneNumber")
		err3 := token.JwtVerify(context)
		_, err4 := fybDatabase.UpdateSingleUserByCondition(db, mp1, mp2)
		result = multierror.Append(result, err1, err2, err3, err4)

		code := 200
		if result.ErrorOrNil() == nil {
			context.JSON(code, gin.H{
				"code":    code,
				"message": "修改成功",
			})
		} else {
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
