/*
 * @Author: 思源 syrun528@126.com
 * @Date: 2025-10-19 10:31:19
 * @LastEditors: 思源 syrun528@126.com
 * @LastEditTime: 2025-10-19 10:31:29
 * @FilePath: \go_blog\model\appTypes\image_storage.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package appTypes

import "encoding/json"

// Storage 图片存储类型
type Storage int

const (
	Local Storage = iota // 本地
	Qiniu                // 七牛云
)

// MarshalJSON 实现了 json.Marshaler 接口
func (s Storage) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// UnmarshalJSON 实现了 json.Unmarshaler 接口
func (s *Storage) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*s = ToStorage(str)
	return nil
}

// String 方法返回 Storage 的字符串表示
func (s Storage) String() string {
	var str string
	switch s {
	case Local:
		str = "本地"
	case Qiniu:
		str = "七牛云"
	default:
		str = "未知存储"
	}
	return str
}

// ToStorage 函数将字符串转换为 Storage
func ToStorage(str string) Storage {
	switch str {
	case "本地":
		return Local
	case "七牛云":
		return Qiniu
	default:
		return -1
	}
}
