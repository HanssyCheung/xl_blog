package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"xl-go-blog/global"
	"xl-go-blog/models"
	"xl-go-blog/models/res"
	"xl-go-blog/utils"
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  //文件名
	IsSuccess bool   `json:"is_success"` //是否上传成功
	Msg       string `json:"msg"`        //消息
}

//ImageUploadView 上传单个图片，返回图片的url
func (ImagesApi) ImageUploadView(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
	}

	//判断路径是否存在，因为每次都是动态修改的，不存在就创建
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

	//判断大小是否超出
	var resList []FileUploadResponse

	for _, file := range fileList {
		//判断后缀
		fileName := file.Filename
		suffix := filepath.Ext(fileName)
		fmt.Println(suffix)
		if !utils.InList(suffix, global.WhiteImageList) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "非法文件",
			})
			continue
		}
		filePath := path.Join(basePath, file.Filename)
		//判断大小
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超过设定，当前大小为：%.2fMB,设定大小为：%dMB", size, global.Config.Upload.Size), //.2f 保留两位小数的意思
			})
			continue
		}

		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		byteData, err := io.ReadAll(fileObj)
		imageHash := utils.Md5(byteData)
		fmt.Println(imageHash)
		//去数据库里查这个图片是否存在
		var bannerModel models.BannerModel
		err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
		if err == nil {
			//找到了
			resList = append(resList, FileUploadResponse{
				FileName:  bannerModel.Path,
				IsSuccess: false,
				Msg:       "图片已存在",
			})
			continue
		}

		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			continue
		}
		//图片入库
		global.DB.Create(&models.BannerModel{
			Path: filePath,
			Hash: imageHash,
			Name: fileName,
		})
		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "上传成功",
		})

		//fmt.Println(file.Header)
		//fmt.Println(file.Size)
		//fmt.Println(file.Filename)
	}
	res.OkWithData(resList, c)

	//fileHeader, err := c.FormFile("image")
	//if err != nil {
	//	res.FailWithMessage(err.Error(), c)
	//	return
	//}
	//fmt.Println(fileHeader.Header)
	//fmt.Println(fileHeader.Size)
	//fmt.Println(fileHeader.Filename)
}
