/*
 * @Author: ZoranYuan syrun0910@outlook.com
 * @Date: 2025-10-19 19:00:59
 * @LastEditors: ZoranYuan syrun0910@outlook.com
 * @LastEditTime: 2025-10-19 19:01:05
 * @FilePath: \go_blog\model\database\friend_link.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package database

import "ZoranYuan_blog/global"

// FriendLink 友链表
type FriendLink struct {
	global.Model
	Logo        string `json:"logo" gorm:"size:255"` // Logo
	Image       Image  `json:"-" gorm:"foreignKey:Logo;references:URL"`
	Link        string `json:"link"`        // 链接
	Name        string `json:"name"`        // 名称
	Description string `json:"description"` // 描述
}
