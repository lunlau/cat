package log

import (
	"fmt"
	"sample/plugins"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerFactory struct {
	loggerHandlerMap map[string]loggerHandler // logger1Name 2 logger. logger2Name 2 logger
}

func (l *LoggerFactory) PluginParentType() string {
	return "log"
}

func (l *LoggerFactory) Run(loggerID string, decoder plugins.Decoder) error {
	cfg := &LoggerHandlerCfg{}
	decoder.Decode(cfg)
	logger := NewLoggerHandler(cfg)
	// 设置全局日志类型
	if loggerID == "global" {
		defaultLogger = logger
	}
	l.rigisterLogger(loggerID, logger)
	return nil
}

func (l *LoggerFactory) rigisterLogger(loggerName string, handler loggerHandler) {
	if loggerName == "" || handler == nil {
		fmt.Println("[ERROR]RigisterLogger loggerName or handle is nil")
		return
	}
	l.loggerHandlerMap[loggerName] = handler
}

//
func (l *LoggerFactory) Sync() {
	for _, v := range l.loggerHandlerMap {
		// new logger,存储logger
		v.Sync()
	}
}

func init() {
	defaultLogger = NewDefaultLoggerHandler()
	plugins.Register("log", &LoggerFactory{})
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

/*
lumberJackLogger := &lumberjack.Logger{
		Filename:   "./run_test.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
*/
func getLogWriter(cfg *LoggerHandlerCfg) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

type LoggerHandlerCfg struct {
	// Filename is the file to write logs to.  Backup log files will be retained
	// in the same directory.  It uses <processname>-lumberjack.log in
	// os.TempDir() if empty.
	Filename string `json:"filename" yaml:"filename"`
	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `json:"maxsize" yaml:"maxsize"`
	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `json:"maxage" yaml:"maxage"`
	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int `json:"maxbackups" yaml:"maxbackups"`
	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool `json:"localtime" yaml:"localtime"`
	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool `json:"compress" yaml:"compress"`
	size     int64
}
