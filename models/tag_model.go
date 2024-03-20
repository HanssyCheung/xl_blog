package models

const (
	TagArticle = 1
	TagSite    = 2
)

type TagModel struct {
	MODEL
	Title   string         `gorm:"size:16" json:"title"`                 //标签名称
	Article []ArticleModel `gorm:"many2many:article_tag" json:"article"` //关联该标签的文章列表
}
