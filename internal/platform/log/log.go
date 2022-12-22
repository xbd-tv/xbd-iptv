package log

import (
	"os"

	logrus "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)
}

type Fields map[string]interface{}

func WithFields(fields Fields) *logrus.Entry {
	f := logrus.Fields{}
	for k, v := range fields {
		f[k] = v
	}

	return logrus.WithFields(f)
}
func Trace(args ...interface{}) {
	logrus.Trace(args...)
}
func Debug(args ...interface{}) {
	logrus.Debug(args...)
}
func Info(args ...interface{}) {
	logrus.Info(args...)
}
func Warn(args ...interface{}) {
	logrus.Warn(args...)
}
func Error(args ...interface{}) {
	logrus.Error(args...)
}
func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}
func Panic(args ...interface{}) {
	logrus.Panic(args...)
}
