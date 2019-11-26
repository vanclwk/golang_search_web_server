package logging_logics

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go_search_online/src/search_online/logics/config_logics"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	ServiceLogger *zap.Logger
	config *viper.Viper
	level string
)

func init() {
	fmt.Printf("Loading Logging module...\n")
	config = config_logics.GlobalConfig
	level = config.GetString("service.logging.level")
	ServiceLogger = BuildLogger(level, "serviceLog")
}

func BuildLogger(level string, logName string) *zap.Logger{
	var logPath string
	var core zapcore.Core
	var logLevel zapcore.Level
	switch level {
	case "Debug":
		logLevel = zapcore.DebugLevel
	case "Info":
		logLevel = zapcore.InfoLevel
	case "Warn":
		logLevel = zapcore.WarnLevel
	case "Error":
		logLevel = zapcore.ErrorLevel
	}
	// set the log level
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(logLevel)
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:"Time",
		LevelKey:"Level",
		NameKey: "Logger",
		CallerKey: "Caller",
		MessageKey:"Msg",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	// logpath the configuration
	debug := config.GetString("service.mode")
	if debug != "Debug" {
		loggerConfig := "service.logging." + logName
		logPath = config.GetString(loggerConfig)
		hook := lumberjack.Logger{
			Filename:   logPath, //日志文件路径
			MaxSize:    200,     // 每个日志的大小，单位是M
			MaxAge:     7,       // 文件被保存的天数
			Compress:   true,    // 是否压缩
			MaxBackups: 10,      // 保存多少个文件备份
		}
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)),
			atomicLevel,
		)
	}else {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
			atomicLevel,
		)
	}
	caller := zap.AddCaller()
	development := zap.Development()
	filed := zap.Fields(zap.String("service","go_search_online"))
	logger := zap.New(core, caller, development, filed)
	return logger
}