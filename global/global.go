package global

import (
	"gorm.io/gorm"
	"xl-go-blog/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
