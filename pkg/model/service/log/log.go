package log

import (
	log "github.com/sirupsen/logrus"
)

// Error 直ちに対処する必要のない実行時エラー
func Error(msg string) {
	log.WithFields(log.Fields{}).Error(msg)
}

// Fatal 致命的なエラー
func Fatal(err error) {
	log.WithFields(log.Fields{}).Fatal(err)
}

// Info 関心を寄せておきたいもの
func Info(msg string) {
	log.WithFields(log.Fields{}).Info(msg)
}

// Warn エラーではないが例外的なもの
func Warn(err error) {
	log.WithFields(log.Fields{}).Warn(err)
}
