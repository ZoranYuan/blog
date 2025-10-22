package config

// ZapConfig 日志配置（zap）
type ZapConfig struct {
	Level          string `yaml:"level" json:"level"`
	Filename       string `yaml:"filename" json:"filename"`
	MaxSize        int    `yaml:"max_size" json:"max_size"`
	MaxBackups     int    `yaml:"max_backups" json:"max_backups"`
	MaxAge         int    `yaml:"max_age" json:"max_age"`
	IsConsolePrint bool   `yaml:"is_console_print" json:"is_console_print"`
}
