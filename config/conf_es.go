package config

// EsConfig Elasticsearch 配置
type EsConfig struct {
	Url            string `yaml:"url" json:"url"`
	Username       string `yaml:"username" json:"username"`
	Password       string `yaml:"password" json:"password"`
	IsConsolePrint bool   `yaml:"is_console_print" json:"is_console_print"`
}
