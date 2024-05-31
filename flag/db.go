package flag

import (
	"xl-go-blog/global"
	"xl-go-blog/models"
)

//MakeMigrations 自动迁移可以新增但不会删除
func MakeMigrations() {
	//fmt.Println("????")
	var err error
	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{}) //设置连表
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	//生成四张表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.BannerModel{},
			&models.TagModel{},
			&models.MessageModel{},
			&models.AdvertModel{},
			&models.UserModel{},
			&models.CommentModel{},
			&models.ArticleModel{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			&models.FeedBackModel{},
			&models.LoginDataModel{},
		)
	if err != nil {
		global.Log.Error("[error] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[success] 生成数据库表结构成功")
}
