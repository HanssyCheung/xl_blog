package settings_api

import (
	"github.com/gin-gonic/gin"
	"xl-go-blog/global"
	"xl-go-blog/models/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	//c.JSON(200, gin.H{"msg": "xxx"})
	//res.Ok(map[string]string{}, "xxx", c)
	//res.OkWithData(map[string]string{
	//	"id": "xxx",
	//}, c)
	//res.FailWithCode(res.SettingsError, c)
	res.OkWithData(global.Config.SiteInfo, c)
}
