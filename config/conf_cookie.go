package config

// EmailConfig 邮件配置
type CookieConfig struct {
	RefreshTokenName string `yaml:"refresh_token_name" json:"refreshTokenName"`
	AccessTokenName  string `yaml:"access_token_name" json:"accessTokenName"`
}
