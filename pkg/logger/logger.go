package logger

import (
	"os"
	"strconv"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	defaultFormatter *TextFormatter
	Log              logger
)

type logger struct {
	*logrus.Logger
}

const (
	PanicLevel logrus.Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func init() {
	defaultFormatter = &TextFormatter{
		DisableColors:    true,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		QuoteEmptyFields: true,
		ForceFormatting:  true,
		QuoteCharacter:   "",
		Once:             sync.Once{},
	}

	Log = logger{New()}
}

func New() *logrus.Logger {
	l := logrus.New()
	l.SetReportCaller(true)
	l.SetFormatter(defaultFormatter)
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))

	if debug {
		l.SetLevel(DebugLevel)
	} else {
		l.SetLevel(InfoLevel)
	}

	return l
}

func (l logger) New() *logrus.Logger {
	return New()
}
