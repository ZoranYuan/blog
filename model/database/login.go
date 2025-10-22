package database

import "ZoranYuan_blog/global"

type Login struct {
	global.Model
	UserID      uint   `json:"user_id"`
	User        User   `json:"user" gorm:"foreignKey:UserID"` // 通过 UserID 来与 User 表中 ID 对应（默认）
	LoginMethod string `json:"login_method"`
	IP          string `json:"ip"`
	Address     string `json:"adress"`
	OS          string
	DeviceInfo  string `json:"device_info"`
	BrowserInfo string `json:"browser_info"`
	Status      int    `json:"status"` // 登录状态
}
