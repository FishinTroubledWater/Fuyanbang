package main

import (
	fybDatabase "FybBackend/database"
	"FybBackend/routers/v1/frontend/academy"
	"FybBackend/routers/v1/frontend/user/login"
	"FybBackend/routers/v1/frontend/user/selectUsers"
	"FybBackend/routers/v1/frontend/user/settings"
	"FybBackend/routers/v1/frontend/user/userInfo"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	db := fybDatabase.InitDB()
	r := gin.Default()
	academy.SearchByRule(r)
	login.PasswordLogin(r, db)
	selectUsers.SelectUsers(r, db)
	userInfo.BasicUserInfo(r, db)
	settings.Settings(r, db)
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}

}
