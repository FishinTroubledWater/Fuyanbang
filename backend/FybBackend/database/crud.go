package database

import (
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func SearchByRegionLevelType(db *gorm.DB, region string, level string, _type string) (error, []Academy, int64) {
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

func SelectSingleUserByCondition(db *gorm.DB, where map[string]interface{}) (User, int64, error) {
	var count int64 = 0
	var user User
	err := db.Where(where).First(&user).Count(&count).Error
	return user, count, err
}
func SelectAllUserByCondition(db *gorm.DB, where map[string]interface{}) ([]User, int64, error) {
	var count int64 = 0
	var users []User
	err := db.Where(where).Find(&users).Count(&count).Error
	return users, count, err
}
func UpdateSingleUserBy(db *gorm.DB, where map[string]interface{}, update map[string]interface{}) (int64, error) {
	var count int64 = 0
	err := db.Table("user").Where(where).Updates(update).Count(&count).Error
	return count, err
}
