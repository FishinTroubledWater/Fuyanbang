package routers

import (
	"FybBackend/routers/v1/backend/post/modifyPost"
	"FybBackend/routers/v1/backend/post/selectPost"
	"FybBackend/routers/v1/backend/user/login"
	"FybBackend/routers/v1/backend/user/modifyUser"
	"FybBackend/routers/v1/backend/user/selectUser"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitBackend(r *gin.Engine, db *gorm.DB) {
	login.Login(r, db)
	modifyUser.AddUser(r, db)
	modifyUser.UpdateUser(r, db)
	modifyUser.DeleteUser(r, db)
	selectUser.SelectUserByAccount(r, db)
	selectUser.SelectUsersByPage(r, db)
	selectPost.SelectPostByPage(r, db)
	modifyPost.DeletePost(r, db)
}
