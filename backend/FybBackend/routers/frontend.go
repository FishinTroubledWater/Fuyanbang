package routers

import (
	"FybBackend/routers/v1/frontend/academy"
	"FybBackend/routers/v1/frontend/circle"
	"FybBackend/routers/v1/frontend/major"
	"FybBackend/routers/v1/frontend/news"
	"FybBackend/routers/v1/frontend/user/login"
	"FybBackend/routers/v1/frontend/user/register"
	"FybBackend/routers/v1/frontend/user/selectUsers"
	"FybBackend/routers/v1/frontend/user/userInfo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitFrontend(r *gin.Engine, db *gorm.DB) {
	//user
	selectUsers.SelectUsers(r, db)
	userInfo.BasicUserInfo(r, db)
	userInfo.Settings(r, db)
	userInfo.ResetPassword(r, db)
	userInfo.GetPassword(r, db)
	userInfo.AccountSecurity(r, db)
	userInfo.AddFeedback(r, db)
	userInfo.MyPoss(r, db)
	register.UserRegister(r, db)
	login.PasswordLogin(r, db)

	//news
	news.NewsInfo(r, db)
	news.NewsDetail(r, db)

	//major
	major.SearchByRule(r, db)
	major.SearchByName(r, db)
	major.SelectMajorByCode(r, db)

	//academy
	academy.SearchByName(r)
	academy.SelectAcademyByCode(r)
	academy.SelectScoreByTypeFirstSecondLevel(r)
	academy.SearchByRule(r)

	//circle
	circle.SearchByName(r)
	circle.SearchNewInfoComment(r)
	circle.SearchNewInfoDetails(r)
	circle.SearchPostComments(r)
	circle.SearchNewQue(r)
}
