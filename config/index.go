/*
 * @Author: 思源 syrun528@126.com
 * @Date: 2025-10-19 08:34:41
 * @LastEditors: 思源 syrun528@126.com
 * @LastEditTime: 2025-10-19 10:28:50
 * @FilePath: \go_blog\config\index.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

// Config 总配置结构体，对应整个配置文件
type Config struct {
	Captcha CaptchaConfig `yaml:"captcha" json:"captcha"`
	Email   EmailConfig   `yaml:"email" json:"email"`
	Es      EsConfig      `yaml:"es" json:"es"`
	Gaode   GaodeConfig   `yaml:"gaode" json:"gaode"`
	Jwt     JwtConfig     `yaml:"jwt" json:"jwt"`
	Mysql   MysqlConfig   `yaml:"mysql" json:"mysql"`
	Qiniu   QiniuConfig   `yaml:"qiniu" json:"qiniu"`
	Qq      QqConfig      `yaml:"qq" json:"qq"`
	Redis   RedisConfig   `yaml:"redis" json:"redis"`
	System  SystemConfig  `yaml:"system" json:"system"`
	Upload  UploadConfig  `yaml:"upload" json:"upload"`
	Website WebsiteConfig `yaml:"website" json:"website"`
	Zap     ZapConfig     `yaml:"zap" json:"zap"`
	Cookie  CookieConfig  `yaml:"cookie" json:"cookie"`
}
