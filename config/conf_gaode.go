package config

// GaodeConfig 高德地图配置
type GaodeConfig struct {
	Enable bool   `yaml:"enable" json:"enable"`
	Key    string `yaml:"key" json:"key"`
}
