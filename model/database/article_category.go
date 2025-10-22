/*
 * @Author: ZoranYuan syrun0910@outlook.com
 * @Date: 2025-10-19 18:56:40
 * @LastEditors: ZoranYuan syrun0910@outlook.com
 * @LastEditTime: 2025-10-19 18:57:17
 * @FilePath: \go_blog\model\database\article_category.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package database

// ArticleCategory 文章类别表
type ArticleCategory struct {
	Category string `json:"category" gorm:"primaryKey"` // 类别
	Number   int    `json:"number"`                     // 统计数量
}
