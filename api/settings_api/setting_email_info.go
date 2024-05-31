package settings_api

import (
	"github.com/gin-gonic/gin"
	"xl-go-blog/global"
	"xl-go-blog/models/res"
)

func (SettingsApi) SettingsEmailInfoView(c *gin.Context) {

	res.OkWithData(global.Config.Email, c)
}
