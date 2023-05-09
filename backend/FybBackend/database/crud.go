package database

import (
	"errors"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func SearchAcademyByRegionLevelType(db *gorm.DB, region string, level string, _type string) (error, []Academy, int64) {
	var academy []Academy
	var result error
	err := db.Table("academy").Where("region=? AND level=? AND type=?", region, level, _type).Find(&academy).Error
	if err != nil {
		result = multierror.Append(result, err)
	}
	var count int64
	err2 := db.Table("academy").Where("region=? AND level=? AND type=?", region, level, _type).Count(&count).Error
	if err2 != nil {
		result = multierror.Append(result, err2)
	}
	return result, academy, count
}

func SearchAcademyByName(db *gorm.DB, name string) (error, []Academy, int64) {
	var academy []Academy
	var result error
	err := db.Table("academy").Where("name=?", name).Find(&academy).Error
	if err != nil {
		result = multierror.Append(result, err)
	}
	var count int64
	err2 := db.Table("academy").Where("name=?", name).Find(&academy).Count(&count).Error
	if err2 != nil {
		result = multierror.Append(result, err2)
	}
	return result, academy, count
}

func SearchAcademyByCode(db *gorm.DB, code string) (error, Academy, int64) {
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

func SelectMajorByCondition(db *gorm.DB, where map[string]interface{}) ([]Major, int64, error) {
	var majors []Major
	var count int64 = 0
	err := db.Table("major").Where(where).Find(&majors).Count(&count).Error
	if count == 0 && err == nil {
		return majors, count, nil
	}
	return majors, count, err
}

func SearchMajorByName(db *gorm.DB, name string) (error, []Academy, int64) {
	var academy []Academy
	var result error
	err := db.Table("academy").Where("name=?", name).Find(&academy).Error
	if err != nil {
		result = multierror.Append(result, err)
	}
	var count int64
	err2 := db.Table("academy").Where("name=?", name).Find(&academy).Count(&count).Error
	if err2 != nil {
		result = multierror.Append(result, err2)
	}
	return result, academy, count
}

func SearchMajorByCode(db *gorm.DB, code string) (error, Academy, int64) {
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
		return user, 0, errors.New("查询的记录不存在")
	}
	return user, count, err
}
func SelectAllUserByCondition(db *gorm.DB, where map[string]interface{}) ([]User, int64, error) {
	var count int64 = 0
	var users []User
	err := db.Where(where).Find(&users).Count(&count).Error
	if count == 0 {
		return users, 0, nil
	}
	return users, count, err
}

func SelectAllUserByPage(db *gorm.DB, pageNum int64, pageSize int64) ([]User, int64, error) {
	var count int64 = 0
	var users []User
	db.Table("user").Count(&count)
	err := db.Table("user").Where(" id >= ? and id <= ?", (pageNum-1)*pageSize+1, pageNum*pageSize).Find(&users).Error
	if count == 0 && err == nil {
		return users, 0, nil
	}
	return users, count, err
}

func AddUser(db *gorm.DB, values map[string]interface{}) (int64, error) {
	var count int64 = 0
	err := db.Table("user").Where("account = ? ", values["account"]).Count(&count).Error
	if count > 0 && err == nil {
		return 0, errors.New("用户已存在")
	}
	err = db.Table("user").Create(values).Count(&count).Error
	return count, err
}

func DeleteUser(db *gorm.DB, where map[string]interface{}) (int64, error) {
	var count int64 = 0
	var user User
	err := db.Table("user").Where("account = ? ", where["account"]).Count(&count).Find(&user).Error
	if count == 0 && err == nil {
		return 0, errors.New("用户不存在")
	}
	err = db.Table("user").Delete(&user).Error
	return count, err
}
func SelectSingleAdminByCondition(db *gorm.DB, where map[string]interface{}) (Admin, int64, error) {
	var count int64 = 0
	var admin Admin
	err := db.Where(where).First(&admin).Count(&count).Error
	return admin, count, err
}
func UpdateSingleAdminByCondition(db *gorm.DB, where map[string]interface{}, update map[string]interface{}) (int64, error) {
	var count int64 = 0
	err := db.Table("admin").Where(where).Updates(update).Count(&count).Error
	return count, err
}

func UpdateSingleUserByCondition(db *gorm.DB, where map[string]interface{}, update map[string]interface{}) (int64, error) {
	var count int64 = 0
	err := db.Table("user").Where(where).Count(&count).Error
	if count == 0 && err == nil {
		return 0, errors.New("用户不存在")
	}
	err = db.Table("user").Where(where).Updates(update).Count(&count).Error
	return count, err
}

func SearchAllNewInfo(db *gorm.DB) ([]Post, error) {
	var result *multierror.Error
	var posts []Post
	err := db.Table("post").Preload("PostImgs").Find(&posts).Error
	if err != nil {
		result = multierror.Append(result, err)
	}
	return posts, err
}
