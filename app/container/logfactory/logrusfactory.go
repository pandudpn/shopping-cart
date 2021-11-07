package logfactory

import (
	"fmt"
	"os"
	"strings"

	"github.com/pandudpn/shopping-cart/src/utils/logger"
	"github.com/sirupsen/logrus"
)

const (
	layoutDate = "2006-01-02 15:04:05 -07:00"
)

type logrusFactory struct{}

type logrusFormatted struct {
	logrus.TextFormatter
}

func (lf *logrusFactory) Build() {
	registerLog()
}

func registerLog() {
	log := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &logrusFormatted{
			logrus.TextFormatter{
				ForceColors:            true,
				FullTimestamp:          true,
				TimestampFormat:        layoutDate,
				DisableLevelTruncation: true,
			},
		},
	}

	logger.SetLoggeer(log)
}

func (f *logrusFormatted) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = 35 // purple
	case logrus.WarnLevel:
		levelColor = 33 // yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}
	return []byte(fmt.Sprintf("[%s] - \x1b[%dm%s\x1b[0m - %s\n", entry.Time.Format(f.TimestampFormat), levelColor, strings.ToUpper(entry.Level.String()), entry.Message)), nil
}
