package myPosts

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"net/http"
	"strconv"
	"time"
)

func SearchNewQue(e *gin.Engine) {
	db := fybDatabase.InitDB()
	e.GET("/v1/frontend/user/myPosts/:authorID", func(context *gin.Context) {
		var result *multierror.Error
		var count int64
		var posts []fybDatabase.Post
		authorId := context.Param("authorID")

		userIdInt64, err := strconv.ParseInt(authorId, 10, 64)

		posts, count, err = fybDatabase.SelectAllPostsByAuthorId(db, userIdInt64)
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

				postMap["id"] = postMap["ID"]
				if publishTime, ok := postMap["PublishTime"].(string); ok {
					t, err := time.Parse(time.RFC3339, publishTime)
					if err == nil {
						postMap["time"] = t.Format("2006.01.02 15:04:05")
					}
				}
				postMap["partID"] = postMap["PartID"]
				postMap["title"] = postMap["Title"]
				postMap["favorite"] = postMap["Favorite"]
				postMap["like"] = postMap["Like"]
				postMap["content"] = postMap["Summary"]

				delete(postMap, "Answer")
				delete(postMap, "Title")
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
				"message": "获取发表过的帖子成功！",
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
