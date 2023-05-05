package main

import (
	"FybBackend/routers/v1/backend/user/login"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initBackend(e *gin.Engine, db *gorm.DB) {
	login.Login(e, db)
}
