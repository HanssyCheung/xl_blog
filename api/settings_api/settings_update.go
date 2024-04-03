package settings_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xl-go-blog/config"
	"xl-go-blog/core"
	"xl-go-blog/global"
	"xl-go-blog/models/res"
)

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	fmt.Println("before", global.Config)
	global.Config.SiteInfo = cr
	fmt.Println("after", global.Config)
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWith(c)
}
