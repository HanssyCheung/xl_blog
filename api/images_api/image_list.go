package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xl-go-blog/global"
	"xl-go-blog/models"
	"xl-go-blog/models/res"
)

type Page struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `json:"sort"`
}

func (ImagesApi) ImageListView(c *gin.Context) {
	var cr Page
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var imageList []models.BannerModel

	count := global.DB.Find(&imageList).RowsAffected
	offset := (cr.Page - 1) * cr.Limit
	if offset < 0 {
		offset = 0
	}

	fmt.Println(count)
	global.DB.Limit(1).Offset(1).Find(&imageList)

	res.OkWithData(gin.H{"count": count, "list": imageList}, c)

	return
}
