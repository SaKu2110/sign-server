package log

import (
	log "github.com/sirupsen/logrus"
)

func Info(msg string) {
	log.WithFields(log.Fields{}).Info(msg)
}

func Warn(err error) {
	log.WithFields(log.Fields{}).Warn(err)
}

func Fatal(err error) {
	log.WithFields(log.Fields{}).Fatal(err)
}
