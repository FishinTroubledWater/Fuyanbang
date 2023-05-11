package userInfo

import (
	fybDatabase "FybBackend/database"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func MyPoss(e *gin.Engine, db *gorm.DB) {
	e.GET("/v1/frontend/user/myPoss", func(context *gin.Context) {
		var result *multierror.Error
		authorID := context.DefaultQuery("authorID", "")
		posts, _, err := fybDatabase.SelectAllPossByUser(db, authorID)
		result = multierror.Append(result, err)
		if result.ErrorOrNil() == nil {

			context.JSON(200, gin.H{
				"code":    200,
				"message": "用户个人信息返回成功",
				"data":    posts,
			})
		} else {
			context.JSON(404, gin.H{
				"code":    404,
				"message": result.Error(),
			})
		}
	})
}
