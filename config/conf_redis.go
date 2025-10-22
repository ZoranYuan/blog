package config

// RedisConfig Redis 配置
type RedisConfig struct {
	Address  string `yaml:"address" json:"address"`
	Password string `yaml:"password" json:"password"`
	Db       int    `yaml:"db" json:"db"`
	Timeout  string `yaml:"timeout" json:"timeout"`
}
