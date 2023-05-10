package routers

import (
	"FybBackend/routers/v1/backend/user/login"
	"FybBackend/routers/v1/backend/user/modify"
	"FybBackend/routers/v1/backend/user/select"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitBackend(r *gin.Engine, db *gorm.DB) {
	login.Login(r, db)
	modify.AddUser(r, db)
	modify.UpdateUser(r, db)
	modify.DeleteUser(r, db)
	_select.SelectUserByAccount(r, db)
	_select.SelectUsersByPage(r, db)

}
