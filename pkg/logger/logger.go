package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"sync"
)

var (
	defaultFormatter *TextFormatter
	Logger           logger
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
		QuoteCharacter:   "",
		Once:             sync.Once{},
	}

	Logger = logger{logrus.New()}

	Logger.SetFormatter(defaultFormatter)

	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))

	if debug {
		logrus.SetLevel(DebugLevel)
	} else {
		logrus.SetLevel(InfoLevel)
	}
}

func New() *logrus.Logger {
	l := logrus.New()
	l.SetFormatter(defaultFormatter)

	return l
}

func (l logger) New() *logrus.Logger {
	return New()
}
