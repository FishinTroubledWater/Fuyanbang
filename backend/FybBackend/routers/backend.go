package routers

import (
	"FybBackend/routers/v1/backend/user/add"
	"FybBackend/routers/v1/backend/user/login"
	"FybBackend/routers/v1/backend/user/select"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitBackend(r *gin.Engine, db *gorm.DB) {
	login.Login(r, db)
	_select.SelectUsersByPage(r, db)
	add.AddUser(r, db)
	_select.SelectUserByAccount(r, db)
}
