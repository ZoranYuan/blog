package database

import "ZoranYuan_blog/global"

// JwtBlacklist JWT 黑名单表
type JwtBlacklist struct {
	global.Model
	Jwt string `json:"jwt" gorm:"type:text"` // Jwt
}
