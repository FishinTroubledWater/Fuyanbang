package login

import (
	fybDatabase "FybBackend/database"
	"FybBackend/routers/v1/backend/token"
	"FybBackend/routers/v1/exceptionHandler"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func Login(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/backend/login", func(context *gin.Context) {
		var result *multierror.Error
		mp1 := make(map[string]interface{})
		mp2 := make(map[string]interface{})
		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp1)

		var tokenStr string
		var err3, err4 error
		admin, count1, err1 := fybDatabase.SelectSingleAdminByCondition(db, mp1)
		if count1 != 0 {
			tokenStr, err3 = token.GenerateToken(admin)
			mp2["token"] = tokenStr
			_, err4 = fybDatabase.UpdateSingleAdminByCondition(db, mp1, mp2)
		}

		result = multierror.Append(result, err1, err2, err3, err4)
		//bcrypt.CompareHashAndPassword([]byte(m.HashedPassword), []byte(inputPassword))

		code, msg := exceptionHandler.Handle(result)
		if code == 200 {
			context.JSON(code, gin.H{
				"code":    code,
				"message": "登录成功",
				"account": admin.Account,
				"token":   tokenStr,
			})
		} else {
			context.JSON(code, gin.H{
				"code":    code,
				"message": msg,
			})
		}
	})
}
