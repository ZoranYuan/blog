package config

// QiniuConfig 七牛云配置
type QiniuConfig struct {
	Zone          string `yaml:"zone" json:"zone"`
	Bucket        string `yaml:"bucket" json:"bucket"`
	ImgPath       string `yaml:"img_path" json:"img_path"`
	AccessKey     string `yaml:"access_key" json:"access_key"`
	SecretKey     string `yaml:"secret_key" json:"secret_key"`
	UseHttps      bool   `yaml:"use_https" json:"use_https"`
	UseCdnDomains bool   `yaml:"use_cdn_domains" json:"use_cdn_domains"`
}
