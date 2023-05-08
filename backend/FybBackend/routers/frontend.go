package routers

import (
	"FybBackend/routers/v1/frontend/academy"
	"FybBackend/routers/v1/frontend/circle"
	"FybBackend/routers/v1/frontend/news/newsInfo"
	"FybBackend/routers/v1/frontend/user/login"
	"FybBackend/routers/v1/frontend/user/register"
	"FybBackend/routers/v1/frontend/user/selectUsers"
	"FybBackend/routers/v1/frontend/user/settings"
	"FybBackend/routers/v1/frontend/user/userInfo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitFrontend(r *gin.Engine, db *gorm.DB) {
	academy.SearchByRule(r)
	login.PasswordLogin(r, db)
	selectUsers.SelectUsers(r, db)
	userInfo.BasicUserInfo(r, db)
	settings.Settings(r, db)
	newsInfo.NewsInfo(r, db)
	academy.SearchByName(r)
	academy.SelectAcademyByCode(r)
	academy.SelectScoreByTypeFirstSecondLevel(r)
	register.UserRegister(r, db)
	circle.SearchByName(r)
}
