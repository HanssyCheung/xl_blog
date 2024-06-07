package models

import "time"

type MODEL struct {
	ID       uint      `gorm:"primarykey" json:"id"`            //主键ID
	CreateAt time.Time `gorm:"autoCreateTime" json:"create_at"` //创建时间
	updateAt time.Time `gorm:"autoUpdateTime" json:"update_at"` //更新时间
}
