package config

import (
	"ZoranYuan_blog/model/appTypes"
	"fmt"
	"strings"
)

// SystemConfig 系统基础配置
type SystemConfig struct {
	Host           string `yaml:"host" json:"host"`
	Port           int    `yaml:"port" json:"port"`
	Env            string `yaml:"env" json:"env"`
	RouterPrefix   string `yaml:"router_prefix" json:"router_prefix"`
	UseMultipoint  bool   `yaml:"use_multipoint" json:"use_multipoint"`
	SessionsSecret string `yaml:"sessions_secret" json:"sessions_secret"`
	OssType        string `yaml:"oss_type" json:"oss_type"`
}

func (s SystemConfig) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func (s SystemConfig) Storage() appTypes.Storage {
	switch strings.ToLower(s.OssType) {
	case "local", "Local":
		return appTypes.Local
	case "qiniu", "Qiniu":
		return appTypes.Qiniu
	default:
		return appTypes.Local
	}
}
