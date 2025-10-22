package config

// JwtConfig JWT 配置
type JwtConfig struct {
	AccessTokenSecret      string `yaml:"access_token_secret" json:"access_token_secret"`
	RefreshTokenSecret     string `yaml:"refresh_token_secret" json:"refresh_token_secret"`
	AccessTokenExpiryTime  string `yaml:"access_token_expiry_time" json:"access_token_expiry_time"`
	RefreshTokenExpiryTime string `yaml:"refresh_token_expiry_time" json:"refresh_token_expiry_time"`
	Issuer                 string `yaml:"issuer" json:"issuer"`
}
