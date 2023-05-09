package login

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

func Login(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/backend/login", func(context *gin.Context) {
		var result *multierror.Error
		mp1 := make(map[string]interface{})
		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp1)
		admin, _, err1 := fybDatabase.SelectSingleAdminByCondition(db, mp1)

		result = multierror.Append(result, err1, err2)
		//bcrypt.CompareHashAndPassword([]byte(m.HashedPassword), []byte(inputPassword))
		code := 200
		if result.ErrorOrNil() == nil {
			mp2 := make(map[string]interface{})
			token, _ := token.GenerateToken(admin.Account, admin.PhoneNumber)
			mp2["token"] = token
			fybDatabase.UpdateSingleAdminByCondition(db, mp1, mp2)
			context.JSON(code, gin.H{
				"code":    code,
				"message": "frontend login success!",
				"data": responseItem{
					admin.ID,
					admin.Account,
					token,
				},
			})
		} else {
			code = 404
			context.JSON(code, gin.H{
				"message": code,
				"msg":     result.Error(),
			})
		}
	})
}
