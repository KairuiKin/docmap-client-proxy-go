package common

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()

	// 获取程序的路径
	execPath, err := os.Executable()
	if err != nil {
		logrus.Fatalf("无法获取可执行文件路径: %v", err)
	}
	execDir := filepath.Dir(execPath)

	// 构建 log 目录路径
	logDir := filepath.Join(execDir, "log")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		logrus.Fatalf("无法创建日志目录: %v", err)
	}

	// 构建日志文件的路径
	logFilePath := filepath.Join(logDir, "DocmapProxy.log")
	// 创建一个多输出流
	multiWriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   logFilePath, // 日志文件路径
		MaxSize:    10,          // 最大日志文件大小，单位 MB
		MaxBackups: 3,           // 保留最多的备份文件数
		MaxAge:     7,           // 保留日志的最大天数
		Compress:   true,        // 是否压缩备份日志
	})

	// 设置 lumberjack 来处理日志文件轮转
	Logger.SetOutput(multiWriter)

	// 设置日志级别
	Logger.SetLevel(logrus.TraceLevel)

	// 设置日志格式：文本格式并自定义字段
	Logger.SetFormatter(&logrus.TextFormatter{
		// 设置日志时间的格式
		TimestampFormat: "2006-01-02 15:04:05",
		// 打印日志时显示日志级别
		FullTimestamp: true,
		// 为了让日志的级别在输出时有颜色（例如 INFO 会变为绿色）
		ForceColors: true,
		// 将日志级别写在左侧
		PadLevelText: true,
		// 输出格式可以有多种选择，设置为 logrus.TextFormatter
	})
}
