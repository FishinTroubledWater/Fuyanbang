package circle

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"net/http"
)

func SearchByName(e *gin.Engine) {
	db := fybDatabase.InitDB()
	e.GET("/v1/frontend/circle/newinfo", func(context *gin.Context) {
		var result *multierror.Error

		posts, err := fybDatabase.SearchAllNewInfo(db)
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

			authorname := postMap["Author"].(map[string]interface{})["NickName"]
			postMap["name"] = authorname
			postMap["time"] = postMap["PublishTime"]

			postMap["postId"] = postMap["ID"]
			postMap["isImage"] = "false"
			postMap["summary"] = postMap["Summary"]
			postMap["icon"] = postMap["Author"].(map[string]interface{})["AvatarUrl"]
			delete(postMap, "PublishTime")
			delete(postMap, "Part")
			delete(postMap, "Summary")
			delete(postMap, "AuthorID")
			delete(postMap, "Author")
			delete(postMap, "ID")
			delete(postMap, "State")
			delete(postMap, "Favorite")
			delete(postMap, "Like")
			delete(postMap, "PartID")

			responseBody = append(responseBody, postMap)
		}

		if result.ErrorOrNil() == nil {
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

func mapValuesToArray(m []interface{}) []string {
	var result []string
	for _, item := range m {
		// 将interface{}转换为map[string]interface{}类型
		value, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		// 获取Img字段的值
		img, ok := value["Img"].(string)
		if !ok {
			continue
		}
		// 将Img字段的值添加到结果数组中
		result = append(result, img)
	}
	return result
}
