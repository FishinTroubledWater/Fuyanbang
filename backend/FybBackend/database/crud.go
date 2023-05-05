package database

import (
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

func SelectSingleNewsByCondition(db *gorm.DB, where map[string]interface{}) (News, int64, error) {
	var count int64 = 0
	var news News
	err := db.Where(where).First(&news).Count(&count).Error
	return news, count, err
}

func SelectAllNewsByCondition(db *gorm.DB, where map[string]interface{}) ([]News, int64, error) {
	var count int64 = 0
	var newses []News
	err := db.Where(where).Find(&newses).Count(&count).Error
	return newses, count, err
}

func SelectSingleUserByCondition(db *gorm.DB, where map[string]interface{}) (User, int64, error) {
	var count int64 = 0
	var user User
	err := db.Where(where).First(&user).Count(&count).Error
	return user, count, err
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
func SelectAllUserByCondition(db *gorm.DB, where map[string]interface{}) ([]User, int64, error) {
	var count int64 = 0
	var users []User
	err := db.Where(where).Find(&users).Count(&count).Error
	return users, count, err
}
func UpdateSingleUserByCondition(db *gorm.DB, where map[string]interface{}, update map[string]interface{}) (int64, error) {
	var count int64 = 0
	err := db.Table("user").Where(where).Updates(update).Count(&count).Error
	return count, err
}
