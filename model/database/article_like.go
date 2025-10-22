package database

import "ZoranYuan_blog/global"

// ArticleLike 文章收藏表
type ArticleLike struct {
	global.Model
	ArticleID string `json:"article_id"` // 文章 ID
	UserID    uint   `json:"user_id"`    // 用户 ID
	User      User   `json:"-" gorm:"foreignKey:UserID"`
}
