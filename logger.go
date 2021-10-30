package logrusPlusPlus

import (
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var loggerInstance *LogrusPlusPlus

var loggerOnce sync.Once

type LogrusPlusPlus struct {
	logger *log.Entry
}

func Initialize(appName string, config LoggingConfig, fields Fields) {

	var envConfig LoggingConfig
	cleanenv.ReadEnv(&envConfig)

	if appName == "" {
		panic("App name cannot be empty")
	}

	if config.LogFileActive {
		envConfig.LogFileActive = config.LogFileActive
	}

	if config.LogFilePath != "" {
		envConfig.LogFilePath = config.LogFilePath
	}

	envConfig.LogFilePath += appName + ".log"

	if config.LogFileMaxAgeDay > 0 {
		envConfig.LogFileMaxAgeDay = config.LogFileMaxAgeDay
	}

	if config.LogFileMaxBackups > 0 {
		envConfig.LogFileMaxBackups = config.LogFileMaxBackups
	}

	if config.LogLevel > 0 {
		envConfig.LogLevel = config.LogLevel
	}

	if config.LogPrettyPrintActive {
		envConfig.LogPrettyPrintActive = config.LogPrettyPrintActive
	}

	if config.LogStdoutActive {
		envConfig.LogStdoutActive = config.LogStdoutActive
	}
	loggerOnce.Do(func() {
		loggerInstance = NewLogrusPlusPlusf(envConfig, fields)
	})
}

func Logger() *LogrusPlusPlus {
	return loggerInstance
}

func NewLogrusPlusPlusf(config LoggingConfig, fields Fields) *LogrusPlusPlus {
	logger := NewLogrusPlusPlus(config)
	logger.setDefaultLogField(fields)
	return logger
}

func NewLogrusPlusPlus(config LoggingConfig) *LogrusPlusPlus {
	logger := log.New()
	logLevel := config.LogLevel
	logger.SetFormatter(&log.JSONFormatter{
		PrettyPrint:     config.LogPrettyPrintActive,
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	})

	ioWriters := []io.Writer{}

	if config.LogStdoutActive {
		ioWriters = append(ioWriters, os.Stdout)
	}

	if config.LogFileActive {

		lumberjackLogger := &lumberjack.Logger{
			// Log file abbsolute path, os agnostic
			Filename:   filepath.ToSlash(config.LogFilePath),
			MaxSize:    config.LogFileMaxSizeMb, // MB
			MaxBackups: config.LogFileMaxBackups,
			MaxAge:     config.LogFileMaxAgeDay, // days
			Compress:   config.LogFileCompress,  // disabled by default
		}

		// file, err := os.OpenFile(config.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		// if err == nil {
		ioWriters = append(ioWriters, lumberjackLogger)
	}

	mw := io.MultiWriter(ioWriters...)
	logger.SetOutput(mw)
	logger.SetLevel(log.Level(logLevel))

	return &LogrusPlusPlus{
		logger: logger.WithFields(log.Fields{}),
	}
}

func (l *LogrusPlusPlus) setDefaultLogField(fields Fields) {
	l.logger = l.logger.WithFields(log.Fields(fields))
}

func (l *LogrusPlusPlus) SetDefaultLogField(key string, value interface{}) {
	l.logger = l.logger.WithFields(log.Fields{key: value})
}

func (l *LogrusPlusPlus) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *LogrusPlusPlus) Debugp(fields Fields, args ...interface{}) {
	l.logger.WithFields(log.Fields(fields)).Debug(args...)
}

func (l *LogrusPlusPlus) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *LogrusPlusPlus) Trace(args ...interface{}) {
	l.logger.Trace(args...)
}

func (l *LogrusPlusPlus) Tracep(fields Fields, args ...interface{}) {
	l.logger.WithFields(log.Fields(fields)).Trace(args...)
}

func (l *LogrusPlusPlus) Tracef(format string, args ...interface{}) {
	l.logger.Tracef(format, args...)
}

func (l *LogrusPlusPlus) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *LogrusPlusPlus) Infop(fields Fields, args ...interface{}) {
	l.logger.WithFields(log.Fields(fields)).Info(args...)
}

func (l *LogrusPlusPlus) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *LogrusPlusPlus) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *LogrusPlusPlus) Warnp(fields Fields, args ...interface{}) {
	l.logger.WithFields(log.Fields(fields)).Warn(args...)
}

func (l *LogrusPlusPlus) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *LogrusPlusPlus) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *LogrusPlusPlus) Errorp(fields Fields, args ...interface{}) {
	l.logger.WithFields(log.Fields(fields)).Error(args...)
}

func (l *LogrusPlusPlus) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *LogrusPlusPlus) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *LogrusPlusPlus) Fatalp(fields Fields, args ...interface{}) {
	l.logger.WithFields(log.Fields(fields)).Fatal(args...)
}

func (l *LogrusPlusPlus) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *LogrusPlusPlus) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *LogrusPlusPlus) Panicp(fields Fields, args ...interface{}) {
	l.logger.WithFields(log.Fields(fields)).Panic(args...)
}

func (l *LogrusPlusPlus) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}
