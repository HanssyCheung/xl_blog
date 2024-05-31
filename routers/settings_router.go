package routers

import (
	"xl-go-blog/api"
)

func (rg RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	rg.GET("settings/:name", settingsApi.SettingsInfoView)
	rg.PUT("settings/:name", settingsApi.SettingsInfoUpdateView)
	rg.GET("settings_email", settingsApi.SettingsEmailInfoView)
	rg.PUT("settings_email", settingsApi.SettingsEmailInfoUpdateView)
}
