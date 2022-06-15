package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var root *logger
var rootMutex sync.Mutex

const serviceName = "study"
const brandName = "test"
const protocol = "HTTP"

type Logger interface {
	Trace(ctx ...interface{})
	Debug(ctx ...interface{})
	Info(ctx ...interface{})
	Warn(ctx ...interface{})
	Error(ctx ...interface{})

	Tracef(format string, ctx ...interface{})
	Debugf(format string, ctx ...interface{})
	Infof(format string, ctx ...interface{})
	Warnf(format string, ctx ...interface{})
	Errorf(format string, args ...interface{})

	WithField(key string, value interface{}) Logger
}

func Root() Logger {
	return root
}

func Trace(args ...interface{}) {
	root.Trace(args...)
}
func Debug(args ...interface{}) {
	root.Debug(args...)
}
func Info(args ...interface{}) {
	root.Info(args...)
}
func Warn(args ...interface{}) {
	root.Warn(args...)
}
func Error(args ...interface{}) {
	root.Error(args...)
}

func Tracef(format string, args ...interface{}) {
	root.Tracef(format, args...)
}
func Debugf(format string, args ...interface{}) {
	root.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	root.Infof(format, args...)
}
func Warnf(format string, args ...interface{}) {
	root.Warnf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	root.Errorf(format, args...)
}

func WithField(key string, value interface{}) Logger {
	return root.WithField(key, value)
}

func New(port string) {
	formatter := &MaaSReNewLogFormatter{}
	l := logrus.New()
	l.SetFormatter(formatter)
	l.SetLevel(logrus.TraceLevel)
	l.AddHook(new(CodeLineNumberRenewHook))
	if err := formatter.init(); err != nil {
		l.Errorf("formatter.Init() error = %v", err)
	}
	hostname, _ := os.Hostname()
	entry := l.WithFields(
		logrus.Fields{
			LogURL:      fmt.Sprintf("%s:%s", hostname, port),
			LogProtocol: protocol,
		},
	)
	logger := &logger{
		entry: entry,
	}
	root = logger
}

type logger struct {
	serviceName string
	brandName   string
	protocol    string
	entry       *logrus.Entry
}

func (log *logger) Trace(args ...interface{}) {
	log.entry.Traceln(args...)
}

func (log *logger) Debug(args ...interface{}) {
	log.entry.Debugln(args...)
}

func (log *logger) Info(args ...interface{}) {
	log.entry.Infoln(args...)
}

func (log *logger) Warn(args ...interface{}) {
	log.entry.Warnln(args...)
}

func (log *logger) Error(args ...interface{}) {
	log.entry.Errorln(args...)
}

func (log *logger) Tracef(format string, args ...interface{}) {
	log.entry.Tracef(format, args...)
}

func (log *logger) Debugf(format string, args ...interface{}) {
	log.entry.Debugf(format, args...)
}

func (log *logger) Infof(format string, args ...interface{}) {
	log.entry.Infof(format, args...)
}

func (log *logger) Warnf(format string, args ...interface{}) {
	log.entry.Warnf(format, args...)
}

func (log *logger) Errorf(format string, args ...interface{}) {
	log.entry.Errorf(format, args...)
}

func (log *logger) WithField(key string, value interface{}) Logger {
	entry := log.entry.WithField(key, value)
	return &logger{
		entry: entry,
	}
}
