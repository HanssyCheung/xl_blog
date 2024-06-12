package common

import (
	"fmt"
	"gorm.io/gorm"
	"xl-go-blog/global"
	"xl-go-blog/models"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}

	count = DB.Select("id").Find(&list).RowsAffected
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}

	fmt.Println(count)
	err = DB.Limit(option.Limit).Offset(offset).Find(&list).Error

	return list, count, err
}
