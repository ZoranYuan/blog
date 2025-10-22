/*
 * @Author: ZoranYuan syrun0910@outlook.com
 * @Date: 2025-10-19 19:01:11
 * @LastEditors: ZoranYuan syrun0910@outlook.com
 * @LastEditTime: 2025-10-19 19:01:18
 * @FilePath: \go_blog\model\database\footer_link.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package database

// FooterLink 页脚链接表
type FooterLink struct {
	Title string `json:"title" gorm:"primaryKey"` // 标题
	Link  string `json:"link"`                    // 链接
}
