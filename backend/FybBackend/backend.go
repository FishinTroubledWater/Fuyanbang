package main

import (
	"FybBackend/routers/v1/backend/user/login"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initBackend(r *gin.Engine, db *gorm.DB) {
	login.Login(r, db)
}
