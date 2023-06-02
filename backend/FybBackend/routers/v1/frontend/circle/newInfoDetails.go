package circle

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"net/http"
	"strconv"
	"time"
)

func SearchNewInfoDetails(e *gin.Engine) {
	db := fybDatabase.InitDB()
	e.GET("/v1/frontend/circle/newinfoDetails/:postId/:userId", func(context *gin.Context) {
		var result *multierror.Error
		var isExistedLikeRecord bool
		var isExistedFavoriteRecord bool

		postId := context.Param("postId")
		userId := context.Param("userId")

		userIdInt64, err := strconv.ParseInt(userId, 10, 64)
		postIdInt64, err := strconv.ParseInt(postId, 10, 64)

		err, posts := fybDatabase.SearchNewInfoDetails(db, postIdInt64)
		if err != nil {
			result = multierror.Append(result, err)
		}

		var responseBody [](map[string]interface{})

		for i := range posts {
			data, err := json.Marshal(&posts[i])
			if err != nil {
				result = multierror.Append(result, err)
			}
			var postMap map[string]interface{}
			err = json.Unmarshal(data, &postMap)
			if err != nil {
				result = multierror.Append(result, err)
			}

			err, isExistedLikeRecord = fybDatabase.IsExistedLikeRecord(db, postIdInt64, userIdInt64)
			if err != nil {
				result = multierror.Append(result, err)
			}

			err, isExistedFavoriteRecord = fybDatabase.IsExistedFavoriteRecord(db, postIdInt64, userIdInt64)
			if err != nil {
				result = multierror.Append(result, err)
			}

			if isExistedLikeRecord == true {
				postMap["isLiked"] = "true"
			} else {
				postMap["isLiked"] = "false"
			}

			if isExistedFavoriteRecord == true {
				postMap["isCollected"] = "true"
			} else {
				postMap["isCollected"] = "false"
			}

			err, postMap["likeNum"] = fybDatabase.GetLikeNumByPostId(db, postIdInt64)
			if err != nil {
				result = multierror.Append(result, err)
			}

			postMap["name"] = postMap["Author"].(map[string]interface{})["NickName"]
			if publishTime, ok := postMap["PublishTime"].(string); ok {
				t, err := time.Parse(time.RFC3339, publishTime)
				if err == nil {
					postMap["time"] = t.Format("2006.01.02 15:04:05")
				}
			}
			postMap["postId"] = postMap["ID"]
			postMap["icon"] = postMap["Author"].(map[string]interface{})["AvatarUrl"]
			postMap["summary"] = postMap["Content"]
			delete(postMap, "Content")
			delete(postMap, "Summary")
			delete(postMap, "Part")
			delete(postMap, "Author")
			delete(postMap, "AuthorID")
			delete(postMap, "Favorite")
			delete(postMap, "State")
			delete(postMap, "PartID")
			delete(postMap, "Like")
			delete(postMap, "PublishTime")
			delete(postMap, "ID")
			delete(postMap, "UserID")

			responseBody = append(responseBody, postMap)
		}

		if result.ErrorOrNil() == nil && len(responseBody) != 0 {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "请求成功",
				"data":    responseBody,
			})
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": result.Error(),
				"data":    responseBody,
			})
		}
	})
}
