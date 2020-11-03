package service

import (
	log "github.com/sirupsen/logrus"
)

func ShowLogInfo(msg string, locate string) {
	log.WithFields(log.Fields{
	  "locate": locate,
	}).Info(msg)
}

func ShowLogWarn(err error, locate string) {
	log.WithFields(log.Fields{
		"locate": locate,
	}).Warn(err)
}

func ShowLogFatal(err error, locate string) {
	log.WithFields(log.Fields{
		"locate": locate,
	}).Fatal(err)
}
