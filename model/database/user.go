/*
 * @Author: ZoranYuan syrun0910@outlook.com
 * @Date: 2025-10-19 17:09:24
 * @LastEditors: ZoranYuan syrun0910@outlook.com
 * @LastEditTime: 2025-10-19 18:05:30
 * @FilePath: \go_blog\model\database\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package database

import (
	"ZoranYuan_blog/global"
	"ZoranYuan_blog/model/appTypes"

	"github.com/gofrs/uuid"
)

// User 用户表
type User struct {
	global.Model
	UUID      uuid.UUID         `json:"uuid" gorm:"type:char(36);unique"`              // uuid
	Username  string            `json:"username"`                                      // 用户名
	Password  string            `json:"-"`                                             // 密码, - 表示在转为 json 的时候忽略这个字段
	Email     string            `json:"email"`                                         // 邮箱
	Openid    string            `json:"openid"`                                        // openid
	Avatar    string            `json:"avatar" gorm:"size:255"`                        // 头像：邮箱注册的头像或 QQ 登录的空间头像
	Address   string            `json:"address"`                                       // 地址
	Signature string            `json:"signature" gorm:"default:'签名是空白的，这位用户似乎比较低调。'"` // 签名
	RoleID    appTypes.RoleID   `json:"role_id"`                                       // 角色 ID
	Register  appTypes.Register `json:"register"`                                      // 注册来源
	Freeze    bool              `json:"freeze"`                                        // 用户是否被冻结
}
