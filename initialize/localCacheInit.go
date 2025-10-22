package initialize

import (
	"ZoranYuan_blog/global"
	"ZoranYuan_blog/utils"

	"github.com/songzhibin97/gkit/cache/local_cache"
)

func IniteLocalCache() (*local_cache.Cache, error) {
	refreshTokenExpiry, _ := utils.ParseDuration(global.Config.Jwt.RefreshTokenExpiryTime)

	// 配置本地缓存过期时间（使用刷新令牌过期时间，方便在远程登录或账户冻结时对 JWT 进行黑名单处理）
	cache := local_cache.NewCache(
		local_cache.SetDefaultExpire(refreshTokenExpiry),
	)

	return &cache, nil
}
