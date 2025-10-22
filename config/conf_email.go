package config

// EmailConfig 邮件配置
type EmailConfig struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	From     string `yaml:"from" json:"from"`
	Nickname string `yaml:"nickname" json:"nickname"`
	Secret   string `yaml:"secret" json:"secret"`
	IsSsl    bool   `yaml:"is_ssl" json:"is_ssl"`
}
