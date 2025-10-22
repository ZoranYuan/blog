package initialize

import (
	"ZoranYuan_blog/global"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() (*zap.Logger, error) {
	zapConfig := global.Config.Zap
	writerSyncer := getWriterSyncer(
		zapConfig.Filename,
		zapConfig.MaxSize,
		zapConfig.MaxBackups,
		zapConfig.MaxAge,
	)

	if zapConfig.IsConsolePrint {
		writerSyncer = zapcore.NewMultiWriteSyncer(writerSyncer, zapcore.AddSync(os.Stdout))
	}

	// 4. 获取日志编码器（定义日志的格式，如 JSON 格式、时间格式等）
	encoder := getEncoder()

	// 5. 设置日志级别（从配置中获取，如 "info"、"debug" 等）
	//    ParseLevel 将字符串级别转换为 zapcore.Level 类型
	logLevel, err := zapcore.ParseLevel(zapConfig.Level)
	if err != nil {
		// 如果配置的级别无效，默认使用 InfoLevel
		logLevel = zapcore.InfoLevel
	}

	// 6. 创建核心日志配置（写入器 + 编码器 + 日志级别）
	core := zapcore.NewCore(encoder, writerSyncer, logLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger, nil

}

func getWriterSyncer(filename string, maxSize, maxBackups, maxAge int) zapcore.WriteSyncer {
	lumberJack := lumberjack.Logger{
		Filename:   filename,   // 日志文件名
		MaxSize:    maxSize,    // 单个日志文件最大大小（MB）
		MaxBackups: maxBackups, // 最大备份数
		MaxAge:     maxAge,     // 日志文件最大保留天数
	}
	// 可以理解为为 zap 添加插件
	return zapcore.AddSync(&lumberJack)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}
