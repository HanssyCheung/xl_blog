package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"xl-go-blog/config"
)

var (
	// WhiteImageList 图片上传白名单
	WhiteImageList = []string{
		".jpg",
		"..png",
		".jpeg",
		".xmb",
		".tif",
		".pjp",
		".ico",
		".svg",
		".svgz",
		".jfif",
		".webp",
		".avif",
	}
)

//用来存放全局变量，保存配置文件，方便在main中调用
//全局变量首字母大写
var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
