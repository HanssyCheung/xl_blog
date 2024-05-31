package settings_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xl-go-blog/global"
	"xl-go-blog/models/res"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

//SettingsInfoView 显示某一项的配置
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	fmt.Printf("c %v\n", *c)
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	fmt.Printf("err %v", err)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	fmt.Printf("cr", cr, cr.Name)
	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)

	}

	//res.OkWithData(global.Config.SiteInfo, c)
}
