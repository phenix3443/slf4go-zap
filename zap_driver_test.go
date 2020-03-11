package slf4go_zap

import (
	"testing"

	slog "github.com/go-eden/slf4go"
	"go.uber.org/zap"
)

var (
	cfg = zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: false,
		// DisableCaller:     true,
		DisableStacktrace: true,
		Encoding:          "console",
		EncoderConfig:     zap.NewDevelopmentEncoderConfig(),
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stdout"},
		InitialFields:     map[string]interface{}{"foo": "bar"},
	}
)

// global logger
// go test --run GlobalLogger
func TestGlobalLogger(t *testing.T) {
	Init(&cfg)
	slog.Debug("global logger")
	slog.Warnf("global logger, warnning: %v", "surrender")
}

// default logger
// go test --run DefaultLogger
func TestDefaultLogger(t *testing.T) {
	Init(&cfg)
	l := slog.GetLogger()
	l.Errorf("default logger name=%s", l.Name())
}

// default logger
// go test --run BindFields
func TestBindFields(t *testing.T) {
	Init(&cfg)
	l := slog.GetLogger()
	l.BindFields(slog.Fields{
		"type": "default",
	})
	l.Errorf("default logger name=%s", l.Name())
}

// new logger use with
// go test --run LoggerGenByWith
func TestLoggerGenByWith(t *testing.T) {
	Init(&cfg)
	l := slog.GetLogger()
	l.Infof("with logger name=%s", l.Name())

	l2 := l.WithFields(slog.Fields{"type": "with1"})
	l2.Infof("with logger name=%s", l2.Name())

	l3 := l.WithFields(slog.Fields{"type": "with2"})
	l3.Infof("with logger name=%s", l3.Name())
}

// logger generated by newLogger function
// go test --run LoggerGenByNew
func TestLoggerGenByNew(t *testing.T) {
	Init(&cfg)

	l1 := slog.NewLogger("NewLogger1")
	l1.BindFields(slog.Fields{"type": "new1"})
	l1.Infof("new logger name=%s", l1.Name())

	l2 := slog.NewLogger("NewLogger2")
	l2.BindFields(slog.Fields{"type": "new2"})
	l2.Infof("new logger name=%s", l2.Name())
}
