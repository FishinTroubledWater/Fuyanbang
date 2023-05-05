package main

import (
	"FybBackend/routers/v1/frontend/academy"
	"FybBackend/routers/v1/frontend/news/newsInfo"
	"FybBackend/routers/v1/frontend/user/login"
	"FybBackend/routers/v1/frontend/user/selectUsers"
	"FybBackend/routers/v1/frontend/user/settings"
	"FybBackend/routers/v1/frontend/user/userInfo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initFrontend(e *gin.Engine, db *gorm.DB) {
	academy.SearchByRule(e)
	login.PasswordLogin(e, db)
	selectUsers.SelectUsers(e, db)
	userInfo.BasicUserInfo(e, db)
	settings.Settings(e, db)
	newsInfo.NewsInfo(e, db)
}
