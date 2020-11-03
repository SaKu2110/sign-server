package service

import (
	log "github.com/sirupsen/logrus"
)

func ShowLogInfo(msg string) {
	log.WithFields(log.Fields{
	}).Info(msg)
}

func ShowLogWarn(err error) {
	log.WithFields(log.Fields{
	}).Warn(err)
}

func ShowLogFatal(err error) {
	log.WithFields(log.Fields{
	}).Fatal(err)
}
