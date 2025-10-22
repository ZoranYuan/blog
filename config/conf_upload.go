package config

// UploadConfig 上传配置
type UploadConfig struct {
	Size int    `yaml:"size" json:"size"`
	Path string `yaml:"path" json:"path"`
}
