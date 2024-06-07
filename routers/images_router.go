package routers

import "xl-go-blog/api"

func (rg RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	rg.POST("images", app.ImageUploadView)
	rg.GET("images", app.ImageListView)
}