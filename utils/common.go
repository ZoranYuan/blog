package utils

import (
	"ZoranYuan_blog/global"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
)

func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d) // 去除字符串两端的空格
	if len(d) == 0 {
		return 0, fmt.Errorf("empty duration string")
	}

	// 定义每个单位及其对应的持续时间值
	unitPattern := map[string]time.Duration{
		"d": time.Hour * 24, // "d" 对应 24 小时
		"h": time.Hour,      // "h" 对应 1 小时
		"m": time.Minute,    // "m" 对应 1 分钟
		"s": time.Second,    // "s" 对应 1 秒
	}

	var totalDuration time.Duration // 总持续时间
	re := regexp.MustCompile(`(\d+)([a-zA-Z]+)`)
	matches := re.FindAllStringSubmatch(d, -1)
	for _, m := range matches {
		v, _ := strconv.Atoi(m[1])
		dur, ok := unitPattern[m[2]]

		if !ok {
			global.Log.Error("unkonwn unit", zap.Any("unit", m[2]))
			os.Exit(1)
		}
		totalDuration += time.Duration(v) * dur
	}

	// 返回总的持续时间
	return totalDuration, nil
}
