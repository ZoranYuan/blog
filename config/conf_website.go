package config

// WebsiteConfig 网站信息配置
type WebsiteConfig struct {
	Logo                 string `yaml:"logo" json:"logo"`
	FullLogo             string `yaml:"full_logo" json:"full_logo"`
	Title                string `yaml:"title" json:"title"`
	Slogan               string `yaml:"slogan" json:"slogan"`
	SloganEn             string `yaml:"slogan_en" json:"slogan_en"`
	Description          string `yaml:"description" json:"description"`
	Version              string `yaml:"version" json:"version"`
	CreatedAt            string `yaml:"created_at" json:"created_at"`
	IcpFiling            string `yaml:"icp_filing" json:"icp_filing"`
	PublicSecurityFiling string `yaml:"public_security_filing" json:"public_security_filing"`
	BilibiliUrl          string `yaml:"bilibili_url" json:"bilibili_url"`
	GiteeUrl             string `yaml:"gitee_url" json:"gitee_url"`
	GithubUrl            string `yaml:"github_url" json:"github_url"`
	Name                 string `yaml:"name" json:"name"`
	Job                  string `yaml:"job" json:"job"`
	Address              string `yaml:"address" json:"address"`
	Email                string `yaml:"email" json:"email"`
	QqImage              string `yaml:"qq_image" json:"qq_image"`
	WechatImage          string `yaml:"wechat_image" json:"wechat_image"`
}
