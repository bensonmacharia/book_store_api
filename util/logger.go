package util

import (
	"os"

	"github.com/sirupsen/logrus"
)

var rlog = logrus.New()

func init() {
	rlog.Out = os.Stdout
	rlog.SetLevel(logrus.InfoLevel)
	rlog.SetFormatter(&logrus.JSONFormatter{})
	rlog.SetOutput(os.Stdout)
	//Log to a file
	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	// 	rlog.Out = file
	// } else {
	// 	log.Fatal("Failed to log to file, using default stderr")
	// }
	rlog.Out = os.Stdout
}

func Logger(source string, url string, status int, message string) {
	rlog.WithFields(logrus.Fields{
		"source": source,
		"url":    url,
		"status": status,
	}).Info(message)
}
