package log

import "github.com/sirupsen/logrus"

type LoggerOptions struct {
	Level  logrus.Level
	Module string
}

type Logger struct {
	logger *logrus.Logger
	fields logrus.Fields
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.logger.WithFields(l.fields).Infof(msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.logger.WithFields(l.fields).Errorf(msg, args...)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.logger.WithFields(l.fields).Warnf(msg, args...)
}

func NewLogger(lo LoggerOptions) *Logger {
	var baseLogger = logrus.New()
	var l = &Logger{
		logger: baseLogger,
		fields: logrus.Fields{
			"module": lo.Module,
		},
	}

	l.logger.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
	}
	l.logger.SetLevel(lo.Level)

	return l
}
