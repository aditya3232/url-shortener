package log

import (
	"io"
	"os"
	"path/filepath"

	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var New = logrus.New()

type LogArgs struct {
	Endpoint      string `json:"endpoint"`
	Status        string `json:"status"`
	FromIpAddress string `json:"from_ip_address"`
}

func init() {
	log := New

	// Set the log file path
	logFilePath := filepath.Join("..", "url-shortener", "log", "log.log")

	// Open the log file
	file, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Warnf("error opening file: %v", err)
	}

	// Set the log output to both stdout and the log file
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)

	// Set the log formatter
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "15:04:05 02-01-2006",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
		DisableHTMLEscape: false,
	}

	// Add the config path for viper
	configPath := filepath.Join("..", "tes_backend_developer_golang_bank_ina_muhammad_aditya", "config")
	viper.AddConfigPath(configPath)
}

func Info(args ...interface{}) {
	New.Info(args...)
}

func Infof(format string, args ...interface{}) {
	New.Infof(format, args...)
}

func Error(args ...interface{}) {
	New.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	New.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	New.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	New.Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	New.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	New.Panicf(format, args...)
}

func Warn(args ...interface{}) {
	New.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	New.Warnf(format, args...)
}

func Debug(args ...interface{}) {
	New.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	New.Debugf(format, args...)
}

func Trace(args ...interface{}) {
	New.Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	New.Tracef(format, args...)
}

func Print(args ...interface{}) {
	New.Print(args...)
}

func Printf(format string, args ...interface{}) {
	New.Printf(format, args...)
}

func Log(level logrus.Level, args ...interface{}) {
	New.Log(level, args...)
}

func Logf(level logrus.Level, format string, args ...interface{}) {
	New.Logf(level, format, args...)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return New.WithFields(fields)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return New.WithField(key, value)
}

func WithError(err error) *logrus.Entry {
	return New.WithError(err)
}
