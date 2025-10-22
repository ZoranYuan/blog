/*
 * @Author: ZoranYuan syrun0910@outlook.com
 * @Date: 2025-10-19 19:00:41
 * @LastEditors: ZoranYuan syrun0910@outlook.com
 * @LastEditTime: 2025-10-19 19:00:48
 * @FilePath: \go_blog\model\database\advertisement.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package database

import "ZoranYuan_blog/global"

// Advertisement 广告表
type Advertisement struct {
	global.Model
	AdImage string `json:"ad_image" gorm:"size:255"` // 图片
	Image   Image  `json:"-" gorm:"foreignKey:AdImage;references:URL"`
	Link    string `json:"link"`    // 链接
	Title   string `json:"title"`   // 标题
	Content string `json:"content"` // 内容
}
