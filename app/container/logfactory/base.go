package logfactory

type logFbInterface interface {
	Build()
}

func GetLogFb() logFbInterface {
	return &logrusFactory{}
}
