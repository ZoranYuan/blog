package config

// QqConfig QQ 第三方登录配置
type QqConfig struct {
	Enable      bool   `yaml:"enable" json:"enable"`
	AppId       string `yaml:"app_id" json:"app_id"`
	AppKey      string `yaml:"app_key" json:"app_key"`
	RedirectUri string `yaml:"redirect_uri" json:"redirect_uri"`
}
