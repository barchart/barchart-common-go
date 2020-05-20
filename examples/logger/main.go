package main

import (
	"sync"

	"github.com/barchart/common-go/pkg/logger"
	"github.com/sirupsen/logrus"
)

// Get customized instance logrus
var log = logger.Log

func main() {
	log.Info("some information")
	log.Error("some error")
	log.Print("the same as Info()")
	log.Warn("some warning")
	log.Debug("you will not see this if DEBUG env is false or not specified")

	log.WithField("hello", "world").Info("you can specify some value for output")

	// Remove filename from output
	log.SetReportCaller(false)
	log.Info("without filename")

	// Create new instance of logger and change formatter
	l := logger.New()
	l.SetReportCaller(false)
	l.SetFormatter(&logrus.JSONFormatter{})

	l.Info("another format of output")

	// Set level of logger
	l.SetLevel(logger.ErrorLevel)
	l.SetFormatter(&logger.TextFormatter{
		DisableColors:    true,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		QuoteEmptyFields: true,
		ForceFormatting:  true,
		QuoteCharacter:   "",
		Once:             sync.Once{},
	})
	l.SetReportCaller(true)
	l.Info("won't appear")
	l.Error("will appear")
}
