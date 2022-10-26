// Small wrapper around a `logrus` Logger.

package logger

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	Name   string
	logger logrus.Logger
}

// NON-FORMATTED

func (logger Logger) Trace(args ...any) {
	logger.logger.WithField("name", logger.Name).Trace(args...)
}

func (logger Logger) Debug(args ...any) {
	logger.logger.WithField("name", logger.Name).Debug(args...)
}

func (logger Logger) Info(args ...any) {
	logger.logger.WithField("name", logger.Name).Info(args...)
}

func (logger Logger) Warn(args ...any) {
	logger.logger.WithField("name", logger.Name).Warn(args...)
}

func (logger Logger) Error(args ...any) {
	logger.logger.WithField("name", logger.Name).Error(args...)
}

func (logger Logger) Fatal(args ...any) {
	logger.logger.WithField("name", logger.Name).Fatal(args...)
}

// FORMATTED

func (logger Logger) Tracef(format string, args ...any) {
	logger.logger.WithField("name", logger.Name).Tracef(format, args...)
}

func (logger Logger) Debugf(format string, args ...any) {
	logger.logger.WithField("name", logger.Name).Debugf(format, args...)
}

func (logger Logger) Infof(format string, args ...any) {
	logger.logger.WithField("name", logger.Name).Infof(format, args...)
}

func (logger Logger) Warnf(format string, args ...any) {
	logger.logger.WithField("name", logger.Name).Warnf(format, args...)
}

func (logger Logger) Errorf(format string, args ...any) {
	logger.logger.WithField("name", logger.Name).Errorf(format, args...)
}

func (logger Logger) Fatalf(format string, args ...any) {
	logger.logger.WithField("name", logger.Name).Fatalf(format, args...)
}

func NewLogger(name string) Logger {
	logger := logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &nested.Formatter{
			ShowFullLevel:   true,
			HideKeys:        true,
			TimestampFormat: "2006-01-02 15:04:05",
			FieldsOrder:     []string{"name"},
		},
	}

	return Logger{
		Name:   name,
		logger: logger,
	}
}
