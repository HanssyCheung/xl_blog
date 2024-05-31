package settings_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xl-go-blog/config"
	"xl-go-blog/core"
	"xl-go-blog/global"
	"xl-go-blog/models/res"
)

//好处是减少了代码，坏处是接口的入参和出参不统一

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	fmt.Printf("cr", cr, cr.Name)
	switch cr.Name {
	case "site":
		var info config.SiteInfo
		c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		} //完成参数绑定
		fmt.Println(info)
		global.Config.SiteInfo = info

		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		var info config.Email
		c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		} //完成参数绑定
		fmt.Println(info)
		global.Config.Email = info
	case "qq":
		var info config.QQ
		c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		} //完成参数绑定
		fmt.Println(info)
		global.Config.QQ = info
	case "qiniu":
		var info config.QiNiu
		c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		} //完成参数绑定
		fmt.Println(info)
		global.Config.QiNiu = info
	case "jwt":
		var info config.Jwt
		c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		} //完成参数绑定
		fmt.Println(info)
		global.Config.Jwt = info
	default:
		res.FailWithMessage("没有对应的配置信息", c)
		return
	}
	core.SetYaml()
	res.OkWith(c)
}
