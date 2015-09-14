package Logger

type Logger interface {
	Trace(msg string, params ...interface{})
	Debug(msg string, params ...interface{})
	Info(msg string, params ...interface{})
	Warn(msg string, params ...interface{})
	Error(msg string, params ...interface{})
	Fatal(msg string, params ...interface{})
}
