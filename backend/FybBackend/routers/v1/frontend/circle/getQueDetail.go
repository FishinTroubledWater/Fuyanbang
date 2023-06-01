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

func SearchQueDetails(e *gin.Engine) {
	db := fybDatabase.InitDB()
	e.GET("/v1/frontend/circle/queDetails/:queId/:userId", func(context *gin.Context) {
		var result *multierror.Error
		var count int64
		var posts []fybDatabase.Post
		var isExistedLikeRecord bool
		var isExistedFavoriteRecord bool
		var isExistedAdoptRecord bool

		queId := context.Param("queId")
		userId := context.Param("userId")

		userIdInt64, err := strconv.ParseInt(userId, 10, 64)
		queIdInt64, err := strconv.ParseInt(queId, 10, 64)

		if err != nil {
			result = multierror.Append(result, err)
		}

		count, posts, err = fybDatabase.SearchQueByQueId(db, queIdInt64)
		if err != nil {
			result = multierror.Append(result, err)
		}

		var responseBody []map[string]interface{}
		if count > 0 {
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

				err, isExistedLikeRecord = fybDatabase.IsExistedLikeRecord(db, queIdInt64, userIdInt64)
				if err != nil {
					result = multierror.Append(result, err)
				}

				err, isExistedFavoriteRecord = fybDatabase.IsExistedFavoriteRecord(db, queIdInt64, userIdInt64)
				if err != nil {
					result = multierror.Append(result, err)
				}

				isExistedAdoptRecord, err = fybDatabase.IsExistedAdoptRecord(db, queIdInt64)
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

				if isExistedAdoptRecord == true {
					postMap["isSolved"] = "true"
				} else {
					postMap["isSolved"] = "false"
				}

				if strconv.Itoa(int(postMap["AuthorID"].(float64))) == userId {
					postMap["isMine"] = true
				} else {
					postMap["isMine"] = false
				}

				err, postMap["likeNum"] = fybDatabase.GetLikeNumByPostId(db, queIdInt64)
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
				postMap["queId"] = postMap["ID"]
				postMap["icon"] = postMap["Author"].(map[string]interface{})["AvatarUrl"]
				postMap["content"] = postMap["Content"]

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
		}

		if result.ErrorOrNil() == nil && count > 0 {
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
