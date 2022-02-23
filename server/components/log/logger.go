package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Level log level
type Level int

// log level const
const (
	LevelNil Level = iota
	LevelTrace
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

var defaultLogger loggerHandler // 系统默认日志打印

type loggerHandler interface {
	// 接口函数
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	// 同步处理
	Sync()
	// 其他处理
}

func NewDefaultLoggerHandler() loggerHandler {
	// 返回默认的zap
	return NewLoggerHandler(&LoggerHandlerCfg{
		Filename:   "./run.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	})
}

func NewLoggerHandler(cfg *LoggerHandlerCfg) loggerHandler {
	writeSyncer := getLogWriter(cfg)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	return &loggerHandlerImpl{
		logger: logger,
	}
}

type loggerHandlerImpl struct {
	logger *zap.Logger
}

func (l *loggerHandlerImpl) Sync() {
	l.logger.Sync()
}

func (l *loggerHandlerImpl) Trace(args ...interface{}) {
	l.logger.Info(fmt.Sprint(args...))
}

// 接口函数
func (l *loggerHandlerImpl) Tracef(format string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, args...))
}
func (l *loggerHandlerImpl) Debug(args ...interface{}) {
	l.logger.Debug(fmt.Sprint(args...))
}
func (l *loggerHandlerImpl) Debugf(format string, args ...interface{}) {
	l.logger.Debug(fmt.Sprintf(format, args...))
}
func (l *loggerHandlerImpl) Info(args ...interface{}) {
	l.logger.Info(fmt.Sprint(args...))
}
func (l *loggerHandlerImpl) Infof(format string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, args...))
}
func (l *loggerHandlerImpl) Warn(args ...interface{}) {
	l.logger.Warn(fmt.Sprint(args...))
}
func (l *loggerHandlerImpl) Warnf(format string, args ...interface{}) {
	l.logger.Warn(fmt.Sprintf(format, args...))
}
func (l *loggerHandlerImpl) Error(args ...interface{}) {
	l.logger.Info(fmt.Sprint(args...))
}
func (l *loggerHandlerImpl) Errorf(format string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, args...))
}
func (l *loggerHandlerImpl) Fatal(args ...interface{}) {
	l.logger.Fatal(fmt.Sprint(args...))
}
func (l *loggerHandlerImpl) Fatalf(format string, args ...interface{}) {
	l.logger.Fatal(fmt.Sprintf(format, args...))
}

// 外部调用默认日志
func Trace(args ...interface{}) {
	defaultLogger.Trace(args)
}
func Tracef(format string, args ...interface{}) {
	defaultLogger.Tracef(format, args)
}
func Debug(args ...interface{}) {
	defaultLogger.Debug(args)
}
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args)
}
func Info(args ...interface{}) {
	defaultLogger.Info(args)
}
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args)
}
func Warn(args ...interface{}) {
	defaultLogger.Trace(args)
}
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args)
}
func Error(args ...interface{}) {
	defaultLogger.Error(args)
}
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args)
}
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args)
}
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args)
}

//
//func ts() {
//	InitLogger()
//	defer logger.Sync()
//	simpleHttpGet("www.google.com")
//	simpleHttpGet("http://www.google.com")
//}
