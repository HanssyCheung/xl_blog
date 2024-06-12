package images_api

import (
	"github.com/gin-gonic/gin"
	"xl-go-blog/models"
	"xl-go-blog/models/res"
	"xl-go-blog/service/common"
)

//ImageListView 图片列表查询页
func (ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := common.ComList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	//res.OkWithData(gin.H{"count": count, "list": imageList}, c)
	res.OkWithList(list, count, c)

	return
}
