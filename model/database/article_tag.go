/*
 * @Author: ZoranYuan syrun0910@outlook.com
 * @Date: 2025-10-19 18:57:51
 * @LastEditors: ZoranYuan syrun0910@outlook.com
 * @LastEditTime: 2025-10-19 18:57:59
 * @FilePath: \go_blog\model\database\article_tag.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package database

// ArticleTag 文章标签表
type ArticleTag struct {
	Tag    string `json:"tag" gorm:"primaryKey"` // 标签
	Number int    `json:"number"`                // 统计数量
}
