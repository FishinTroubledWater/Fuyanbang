package database

import (
	"errors"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
	"time"
)

// Academy

func SearchAcademyByRegionLevelType(db *gorm.DB, where map[string]interface{}) (error, []Academy, int64) {
	var result *multierror.Error
	var academy []Academy
	var count int64
	err := db.Table("academy").Where(where).Find(&academy).Count(&count).Error
	if err != nil {
		result = multierror.Append(result, err)
	}
	return result, academy, count
}

func SearchAcademyByName(db *gorm.DB, name string) (error, []Academy, int64) {
	var academy []Academy
	var result error
	name = "%" + name + "%"
	err := db.Table("academy").Where("name like ?", name).Find(&academy).Error
	if err != nil {
		result = multierror.Append(result, err)
	}
	var count int64
	err2 := db.Table("academy").Where("name like ?", name).Find(&academy).Count(&count).Error
	if err2 != nil {
		result = multierror.Append(result, err2)
	}
	return result, academy, count
}

func SearchAcademyByCode(db *gorm.DB, code string) (error, []Academy, int64) {
	var academy []Academy
	var result error
	code = code + "%"
	err := db.Table("academy").Where("code like ?", code).Find(&academy).Error
	if err != nil {
		result = multierror.Append(result, err)
	}
	var count int64
	err2 := db.Table("academy").Where("code like ?", code).Find(&academy).Count(&count).Error
	if err2 != nil {
		result = multierror.Append(result, err2)
	}
	return result, academy, count
}

// Major

func SelectMajorByCondition(db *gorm.DB, where map[string]interface{}) ([]Major, int64, error) {
	var majors []Major
	var count int64 = 0
	err := db.Table("major").Where(where).Find(&majors).Count(&count).Error
	if count == 0 && err == nil {
		return majors, count, nil
	}
	return majors, count, err
}

func SearchMajorByName(db *gorm.DB, name string) (error, []Major, int64) {
	var majors []Major
	var result error
	name = "%" + name + "%"
	err := db.Table("major").Where("name like ?", name).Find(&majors).Error
	if err != nil {
		result = multierror.Append(result, err)
	}
	var count int64
	err2 := db.Table("major").Where("name like ?", name).Find(&majors).Count(&count).Error
	if err2 != nil {
		result = multierror.Append(result, err2)
	}
	return result, majors, count
}

func SearchMajorByCode(db *gorm.DB, code string) (error, []Major, int64) {
	var majors []Major
	code = code + "%"
	var count int64
	err2 := db.Where("code like ?", code).Find(&majors).Count(&count).Error
	if count == 0 && err2 == nil {
		return nil, majors, 0
	}
	return err2, majors, count
}

func SearchScore(db *gorm.DB, code string) (error, Academy, int64) {
	var academy Academy
	var result error
	err := db.Table("academy").Where("code=?", code).Find(&academy).Error
	if err != nil {
		result = multierror.Append(result, err)
	}
	var count int64
	err2 := db.Table("academy").Where("code=?", code).Find(&academy).Count(&count).Error
	if err2 != nil {
		result = multierror.Append(result, err2)
	}
	return result, academy, count
}

func SearchScoreByTypeFirstSecondLevel(db *gorm.DB, where map[string]interface{}) (error, []string) {
	var result *multierror.Error
	var imgUrl []string
	err := db.Table("major").Where(where).Select("scoreUrl").Find(&imgUrl).Error
	if err != nil {
		result = multierror.Append(result, err)
	}
	return result, imgUrl
}

// Comment

func DeleteComment(db *gorm.DB, where map[string]interface{}) (int64, error) {
	var count int64 = 0
	var comment Comment
	err := db.Where(where).Find(&comment).Count(&count).Error
	if count == 0 && err == nil {
		return 0, errors.New("要删除的记录不存在")
	}
	err = db.Delete(&comment).Error
	return count, err
}

func SelectSingleCommentByCondition(db *gorm.DB, where map[string]interface{}) (Comment, int64, error) {
	var count int64 = 0
	var comment Comment
	err := db.Where(where).First(&comment).Count(&count).Error
	if count == 0 {
		return comment, 0, errors.New("查询的记录不存在")
	}
	return comment, count, err
}
func SelectAllCommentByPage(db *gorm.DB, query string, pageNum int64, pageSize int64) ([]Comment, int64, error) {
	var count int64 = 0
	var comments []Comment
	if query != "" {
		query = query + "%"
		db = db.Table("comment").Joins("inner join user on post.authorID = user.id and user.account like ?", query)
	} else {
		db = db.Table("comment")
	}
	db.Table("post").Count(&count)
	err := db.Preload("Author").Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&comments).Error
	if count == 0 && err == nil {
		return comments, 0, errors.New("查询的记录不存在")
	}
	return comments, count, err
}

func AddComments(db *gorm.DB, values map[string]interface{}) (int64, error) {
	var result *multierror.Error
	mp := make(map[string]interface{})
	mp["ID"] = values["userId"]
	delete(values, "account")
	user, count, err := SelectSingleUserByCondition(db, mp)
	result = multierror.Append(result, err)
	if count == 0 {
		result = multierror.Append(result, errors.New("要插入的记录有误，插入的用户不存在"))
	}
	values["userID"] = user.ID
	delete(mp, "ID")
	if _, ok := values["postId"]; ok {
		mp["ID"] = values["postId"]
		delete(values, "postId")
	} else if _, ok := values["queId"]; ok {
		mp["ID"] = values["queId"]
	}

	post, count, err := SelectSingleQueByCondition(db, mp)
	result = multierror.Append(result, err)
	if count == 0 {
		result = multierror.Append(result, errors.New("要插入的记录有误，插入的帖子不存在"))
	}

	if _, ok := values["postId"]; ok {
		values["content"] = values["comment"]
		delete(values, "comment")
	} else if _, ok := values["queId"]; ok {
		values["content"] = values["answer"]
		delete(values, "answer")
		delete(values, "queId")
	}
	values["targetPost"] = post.ID

	delete(values, "userId")
	t1 := time.Now().Year()
	t2 := time.Now().Month()
	t3 := time.Now().Day()
	t4 := time.Now().Hour()
	t5 := time.Now().Minute()
	t6 := time.Now().Second()
	t7 := time.Now().Nanosecond()
	currentTimeData := time.Date(t1, t2, t3, t4, t5, t6, t7, time.Local) //获取当前时间，返回当前时间Time
	values["publishTime"] = currentTimeData
	values["state"] = 0
	err = db.Table("comment").Create(values).Count(&count).Error
	result = multierror.Append(result, err)
	return count, result
}

func SearchCommentByQueId(db *gorm.DB, queId int64) (int64, []Comment, error) {
	var result *multierror.Error
	var comments []Comment
	var whereMap map[string]interface{} = make(map[string]interface{})
	whereMap["ID"] = queId
	whereMap["partID"] = 2
	_, count, err := SelectSingleQueByCondition(db, whereMap)
	fmt.Println(count)
	if count == 0 {
		result = multierror.Append(result, errors.New("帖子id不是问题！"))
	}
	if err != nil {
		result = multierror.Append(result, err)
	}
	err = db.Preload("Author").Where("targetPost = ? ", queId).Find(&comments).Count(&count).Error
	if count == 0 {
		result = multierror.Append(result, errors.New("找不到该问题！"))
	}
	if err != nil {
		result = multierror.Append(result, err)
	}
	return count, comments, result
}

// Recipe

func DeleteRecipe(db *gorm.DB, where map[string]interface{}) (int64, error) {
	var count int64 = 0
	var recipe Recipe
	err := db.Where(where).Find(&recipe).Count(&count).Error
	if count == 0 && err == nil {
		return 0, errors.New("要删除的记录不存在")
	}
	err = db.Delete(&recipe).Error
	return count, err
}

func SelectSingleRecipeByCondition(db *gorm.DB, where map[string]interface{}) (Recipe, int64, error) {
	var count int64 = 0
	var recipe Recipe
	err := db.Where(where).First(&recipe).Count(&count).Error
	if count == 0 {
		return recipe, 0, errors.New("查询的记录不存在")
	}
	return recipe, count, err
}

func SelectAllRecipeByCondition(db *gorm.DB, where map[string]interface{}) ([]Recipe, int64, error) {
	var count int64 = 0
	var recipes []Recipe
	err := db.Where(where).Find(&recipes).Count(&count).Error
	return recipes, count, err
}

func SelectAllRecipeByPage(db *gorm.DB, query string, pageNum int64, pageSize int64) ([]Recipe, int64, error) {
	var count int64 = 0
	var recipes []Recipe

	if query != "" {
		query = query + "%"
		db = db.Table("recipe").Where("author like ?", query).Count(&count)
	} else {
		db = db.Table("recipe").Count(&count)
	}
	err := db.Order("id asc").Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&recipes).Error
	if count == 0 && err == nil {
		return recipes, 0, errors.New("查询的记录不存在")
	}
	return recipes, count, err
}

// News

func DeleteNews(db *gorm.DB, where map[string]interface{}) (int64, error) {
	var count int64 = 0
	var news News
	err := db.Where(where).Find(&news).Count(&count).Error
	if count == 0 && err == nil {
		return 0, errors.New("要删除的记录不存在")
	}
	err = db.Delete(&news).Error
	return count, err
}

func SelectSingleNewsByCondition(db *gorm.DB, where map[string]interface{}) (News, int64, error) {
	var count int64 = 0
	var news News
	err := db.Where(where).First(&news).Count(&count).Error
	if count == 0 {
		return news, 0, errors.New("查询的记录不存在")
	}
	return news, count, err
}

func SelectAllNewsByCondition(db *gorm.DB, where map[string]interface{}) ([]News, int64, error) {
	var count int64 = 0
	var newses []News
	err := db.Where(where).Find(&newses).Count(&count).Error
	return newses, count, err
}

func SelectAllNewsByPage(db *gorm.DB, query string, pageNum int64, pageSize int64) ([]News, int64, error) {
	var count int64 = 0
	var newses []News

	if query != "" {
		query = query + "%"
		db = db.Table("news").Where("author like ?", query).Count(&count)
	} else {
		db = db.Table("news").Count(&count)
	}
	err := db.Order("id asc").Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&newses).Error
	if count == 0 && err == nil {
		return newses, 0, errors.New("查询的记录不存在")
	}
	return newses, count, err
}

func SearchAllNewInfo(db *gorm.DB) ([]Post, error) {
	var result *multierror.Error
	var posts []Post
	err := db.Preload("Author").Where("partID = ? ", 1).Find(&posts).Error
	if err != nil {
		result = multierror.Append(result, err)
	}
	return posts, err
}

func SearchNewInfoComment(db *gorm.DB) (error, []Comment) {
	var result *multierror.Error
	var comments []Comment
	err := db.Preload("Author").Find(&comments).Error
	if err != nil {
		result = multierror.Append(result, err)
	}
	return result, comments
}

func SearchNewInfoDetails(db *gorm.DB, postId int64) (error, []Post) {
	var result *multierror.Error
	var posts []Post
	err := db.Preload("Author").Where("Id = ? && partID = ? ", postId, 1).Find(&posts).Error
	if err != nil {
		result = multierror.Append(result, err)
	}
	return result, posts
}

func UpdateSingleNewsByCondition(db *gorm.DB, where map[string]interface{}, update map[string]interface{}) (int64, error) {
	var count int64 = 0
	err := db.Table("news").Where(where).Count(&count).Error
	if count == 0 && err == nil {
		return 0, errors.New("要修改的记录不存在")
	}
	err = db.Table("news").Where(where).Updates(update).Count(&count).Error
	return count, err
}

func AddNews(db *gorm.DB, values map[string]interface{}) (int64, error) {
	var count int64 = 0
	err := db.Table("news").Create(values).Count(&count).Error
	return count, err
}

// Post ------------------------------------------------------------

func AddPost(db *gorm.DB, values map[string]interface{}) (int64, error) {
	mp := make(map[string]interface{})
	mp["account"] = values["account"]
	delete(values, "account")
	user, count, _ := SelectSingleUserByCondition(db, mp)
	if count == 0 {
		return 0, errors.New("要插入的记录有误")
	}
	values["authorID"] = user.ID
	values["favorite"] = 0
	values["like"] = 0
	err := db.Table("post").Create(values).Count(&count).Error
	return count, err
}

func DeletePost(db *gorm.DB, where map[string]interface{}) (int64, error) {
	var count int64 = 0
	var post Post
	err := db.Where(where).Find(&post).Count(&count).Error
	if count == 0 && err == nil {
		return 0, errors.New("要删除的记录不存在")
	}
	err = db.Delete(&post).Error
	return count, err
}

func SelectSinglePostByCondition(db *gorm.DB, where map[string]interface{}) (Post, int64, error) {
	var count int64 = 0
	var post Post
	err := db.InnerJoins("Author").InnerJoins("Part").Where(where).First(&post).Count(&count).Error
	if count == 0 {
		return post, 0, errors.New("查询的记录不存在")
	}
	return post, count, err
}

func SelectSingleQueByCondition(db *gorm.DB, where map[string]interface{}) (Post, int64, error) {
	var count int64 = 0
	var post Post
	err := db.Where(where).First(&post).Count(&count).Error
	if count == 0 {
		return post, 0, errors.New("查询的记录不存在")
	}
	return post, count, err
}

func SelectAllPostByCondition(db *gorm.DB, where map[string]interface{}) ([]Post, int64, error) {
	var count int64 = 0
	var posts []Post
	err := db.Preload("Author").InnerJoins("Part").Where(where).Count(&count).Error
	return posts, count, err
}
func SelectAllPostByPage(db *gorm.DB, query string, pageNum int64, pageSize int64) ([]Post, int64, error) {
	var count int64 = 0
	var posts []Post
	var err error
	if query != "" {
		query = query + "%"
		db = db.Table("post").InnerJoins("Author").InnerJoins("Part").
			Where("account like ?", query).Order("state asc, id").Find(&posts).Count(&count)
	} else {
		db = db.Table("post").InnerJoins("Author").InnerJoins("Part").
			Order("state asc, id").Find(&posts).Count(&count)
	}
	err = db.Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&posts).Error
	if count == 0 && err == nil {
		return posts, 0, errors.New("要查询的记录不存在")
	}
	return posts, count, err
}

func UpdateSinglePostByCondition(db *gorm.DB, where map[string]interface{}, update map[string]interface{}) (int64, error) {
	var count int64 = 0
	err := db.Table("post").Where(where).Count(&count).Error
	if count == 0 && err == nil {
		return 0, errors.New("要修改的记录不存在")
	}
	err = db.Table("post").Where(where).Updates(update).Count(&count).Error
	return count, err
}

// User ------------------------------------------------------------

func SelectUserForLogin(db *gorm.DB, account string, password string) (User, int64, error) {
	var count int64 = 0
	var user User
	err := db.Where("account = ?", account).First(&user).Count(&count).Error
	if count == 0 {
		return user, 0, errors.New("账户不存在")
	} else if password != user.Password {
		return user, 0, errors.New("密码错误！")
	}
	return user, count, err
}
func SelectSingleUserByCondition(db *gorm.DB, where map[string]interface{}) (User, int64, error) {
	var count int64 = 0
	var user User
	err := db.Where(where).First(&user).Count(&count).Error
	if count == 0 {
		return user, 0, errors.New("要查询的记录不存在")
	}
	return user, count, err
}
func SelectAllUserByCondition(db *gorm.DB, where map[string]interface{}) ([]User, int64, error) {
	var count int64 = 0
	var users []User
	err := db.Where(where).Find(&users).Count(&count).Error
	if count == 0 {
		return users, 0, errors.New("要查询的记录不存在")
	}
	return users, count, err
}

func SelectAllUserByPage(db *gorm.DB, query string, pageNum int64, pageSize int64) ([]User, int64, error) {
	var count int64 = 0
	var users []User
	if query != "" {
		query = query + "%"
		db = db.Table("user").Where("account like ?", query)
	}
	db.Table("user").Count(&count)
	err := db.Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&users).Error
	if count == 0 && err == nil {
		return users, 0, errors.New("要查询的记录不存在")
	}
	return users, count, nil
}

func AddUser(db *gorm.DB, values map[string]interface{}) (int64, error) {
	var count int64 = 0
	err := db.Table("user").Where("account = ? ", values["account"]).Count(&count).Error
	if count > 0 && err == nil {
		return 0, errors.New("记录已存在")
	}
	err = db.Table("user").Create(values).Count(&count).Error
	return count, err
}

func DeleteUser(db *gorm.DB, where map[string]interface{}) (int64, error) {
	var count int64 = 0
	var user User
	err := db.Where(where).Find(&user).Count(&count).Error
	if count == 0 && err == nil {
		return 0, errors.New("要删除的记录不存在")
	}
	err = db.Delete(&user).Error
	return count, err
}

// Feedback

func AddFeedback(db *gorm.DB, values map[string]interface{}) (int64, error) {
	var count int64 = 0
	err := db.Table("user").Create(values).Count(&count).Error
	return count, err
}

func SelectAllFeedbackByPage(db *gorm.DB, query string, pageNum int64, pageSize int64) ([]Feedback, int64, error) {
	var count int64 = 0
	var feedbacks []Feedback
	if query != "" {
		query = query + "%"
		db = db.Table("feedback").Where("author like ?", query).Count(&count)
	} else {
		db = db.Table("feedback").Count(&count)
	}
	err := db.Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&feedbacks).Error
	if count == 0 && err == nil {
		return feedbacks, 0, errors.New("查询的记录不存在")
	}
	return feedbacks, count, err
}
func SelectAllPossByUser(db *gorm.DB, userID string) ([]Post, int64, error) {
	var count int64 = 0
	var posts []Post
	err := db.Where("authorID = ?", userID).Find(&posts).Error
	if count == 0 && err == nil {
		return posts, 0, errors.New("查询的记录不存在")
	}
	return posts, count, nil
}

func UpdateSingleFeedbackByCondition(db *gorm.DB, where map[string]interface{}, update map[string]interface{}) (int64, error) {
	var count int64 = 0
	err := db.Table("feedback").Where(where).Count(&count).Error
	if count == 0 && err == nil {
		return 0, errors.New("要修改的记录不存在")
	}
	err = db.Table("feedback").Where(where).Updates(update).Count(&count).Error
	return count, err
}

// Admin ------------------------------------------------------------

func SelectSingleAdminByCondition(db *gorm.DB, where map[string]interface{}) (Admin, int64, error) {
	var count int64 = 0
	var admin Admin
	err := db.Where(where).First(&admin).Count(&count).Error
	if count == 0 && err == nil {
		return admin, 0, errors.New("查询的记录不存在")
	}
	return admin, count, err
}
func UpdateSingleAdminByCondition(db *gorm.DB, where map[string]interface{}, values map[string]interface{}) (int64, error) {
	var count int64 = 0
	err := db.Table("admin").Where(where).Updates(values).Count(&count).Error
	return count, err
}

func UpdateSingleUserByCondition(db *gorm.DB, where map[string]interface{}, values map[string]interface{}) (int64, error) {
	var count int64 = 0
	err := db.Table("user").Where(where).Count(&count).Error
	if count == 0 && err == nil {
		return 0, errors.New("要修改的记录不存在")
	}
	err = db.Table("user").Where(where).Updates(values).Count(&count).Error
	return count, err
}

// Que ------------------------------------------------------------

func SearchAllQue(db *gorm.DB, userId int64) (int64, []Post, error) {
	var result *multierror.Error
	var posts []Post
	var count int64
	err := db.Preload("comment").Where("partID = ? && authorId = ? ", 2, userId).Find(&posts).Count(&count).Error
	if count == 0 {
		result = multierror.Append(result, errors.New("找不到该用户！"))
	}
	if err != nil {
		result = multierror.Append(result, err)
	}
	return count, posts, result
}

func SearchQueByQueId(db *gorm.DB, queId int64) (int64, []Post, error) {
	var result *multierror.Error
	var posts []Post
	var count int64
	err := db.Preload("Author").Where("ID = ? ", queId).Find(&posts).Count(&count).Error
	if count == 0 {
		result = multierror.Append(result, errors.New("找不到该问题！"))
	}
	if err != nil {
		result = multierror.Append(result, err)
	}
	return count, posts, result
}
