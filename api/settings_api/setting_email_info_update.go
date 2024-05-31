package settings_api

import (
	"github.com/gin-gonic/gin"
	"xl-go-blog/config"
	"xl-go-blog/core"
	"xl-go-blog/global"
	"xl-go-blog/models/res"
)

//SettingsEmailInfoUpdateView 修改某一项的配置信息 封装到一起的弊端就是接口没法统一了
func (SettingsApi) SettingsEmailInfoUpdateView(c *gin.Context) {
	var cr config.Email
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	global.Config.Email = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWith(c)
}
