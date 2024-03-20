package models

import "xl-go-blog/models/ctype"

type ArticleModel struct {
	MODEL
	Title         string         `gorm:"size:32" json:"title"`                      //文章标题
	Abstract      string         `json:"abstract"`                                  //文章简介
	Content       string         `json:"content"`                                   //文章内容
	LookCount     int            `json:"look_count"`                                //浏览量
	CommentCount  int            `json:"comment_count"`                             //评论量
	DiggCount     int            `json:"diggCount"`                                 //点赞量
	CollectsCount int            `json:"collects_count"`                            //收藏量
	TagModels     []TagModel     `gorm:"many2many:article_tag" json:"tagModels"`    //文章标签
	CommentModels []CommentModel `gorm:"foreignKey:ArticleID" json:"commentModels"` //文章的评论列表
	UserModel     UserModel      `gorm:"foreignKey:UserID" json:"userModel"`        //文章作者
	Category      string         `gorm:size:20 json:"category"`                     //文章分类
	Source        string         `gorm:"size:20" json:"source"`                     //文章来源
	Link          string         `json:"link"`                                      //原文链接
	Banner        BannerModel    `json:"banner"`                                    //文章封面
	BannerID      uint           `json:"banner_id"`                                 //文章封面ID
	NickName      string         `gorm:"size:42" json:"nickName"`                   //发布文章的用户昵称
	CoverPath     string         `json:"cover_path"`                                //文章的封面
	Tags          ctype.Array    `gorm:"type:string;size:64" json:"tags"`           //文章标签
}
