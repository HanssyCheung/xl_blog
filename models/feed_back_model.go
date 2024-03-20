package models

// FeedBackModel 用户反馈表
type FeedBackModel struct {
	MODEL
	Email        string `gorm:"size:64" json:"email"`
	Content      string `gorm:"size:128" json:"content"`
	ApplyContent string `gorm:"size:128" json:"applyContent"` //回复的内容
	IsApply      bool   `json:"is_apply"`                     //是否回复
}
