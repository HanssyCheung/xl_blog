package routers

import (
	"xl-go-blog/api"
)

func (rg RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	rg.GET("settings", settingsApi.SettingsInfoView)
}
