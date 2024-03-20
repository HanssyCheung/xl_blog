package models

import "xl-go-blog/models/ctype"

//UserModel 用户表
type UserModel struct {
	MODEL
	NickName       string           `gorm:"size:42" json:"nick_name"`                                                                      //昵称
	UserName       string           `gorm:"size:36" json:"user_name"`                                                                      //用户名
	Password       string           `gorm:"size:128" json:"password"`                                                                      //密码
	Avatar         uint             `gorm:"size:256" json:"avatar"`                                                                        //头像地址
	Email          string           `gorm:"size:128" json:"email"`                                                                         //邮箱
	Tel            string           `gorm:"size:18" json:"tel"`                                                                            //手机号
	Addr           string           `gorm:"size:64" json:"addr"`                                                                           //地址
	Token          string           `gorm:"size:64" json:"token"`                                                                          //其他平台唯一的id
	IP             string           `gorm:"size:20" json:"ip"`                                                                             //ip地址
	Role           ctype.Role       `gorm:"size:4,default:1" json:"role"`                                                                  //权限 1 管理员 2普通用户 3游客
	SignStatus     ctype.SignStatus `gorm:"type=smallint(6)" json:"signStatus"`                                                            //注册来源
	ArticleModels  []ArticleModel   `gorm:"foreignKey:AuthID" json:"articleModels"`                                                        //发布文章列表
	CollectsModels []ArticleModel   `gorm:"many2many:auth2_collects;joinForeignKey:AuthID;JoinReferences:ArticleID" json:"collectsModels"` //收藏的文章列表
}
