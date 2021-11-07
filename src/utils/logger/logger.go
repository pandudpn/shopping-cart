package logger

// Log adalah package level variable
var Log Logger

// Logger mewaikili interface umum untuk logging function
type Logger interface {
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Warn(args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	Printf(format string, args ...interface{})
	Print(args ...interface{})
}

func SetLoggeer(newLogger Logger) {
	Log = newLogger
}
