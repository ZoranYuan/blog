package config

import (
	"strconv"
	"strings"

	"gorm.io/gorm/logger"
)

// MysqlConfig MySQL 配置
type MysqlConfig struct {
	Host         string `yaml:"host" json:"host"`
	Port         int    `yaml:"port" json:"port"`
	Config       string `yaml:"config" json:"config"`
	DbName       string `yaml:"db_name" json:"db_name"`
	Username     string `yaml:"username" json:"username"`
	Password     string `yaml:"password" json:"password"`
	MaxIdleConns int    `yaml:"max_idle_conns" json:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns" json:"max_open_conns"`
	LogMode      string `yaml:"log_mode" json:"log_mode"`
}

func (m MysqlConfig) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DbName + "?" + m.Config
}

func (m MysqlConfig) LogLevel() logger.LogLevel {
	switch strings.ToLower(m.LogMode) {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}
