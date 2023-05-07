package routers

import (
	"FybBackend/routers/v1/backend/user/addUser"
	"FybBackend/routers/v1/backend/user/login"
	"FybBackend/routers/v1/backend/user/selectUsersByPage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitBackend(r *gin.Engine, db *gorm.DB) {
	login.Login(r, db)
	selectUsersByPage.SelectUsersByPage(r, db)
	addUser.AddUser(r, db)
}
