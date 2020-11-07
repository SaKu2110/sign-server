package log

import (
	log "github.com/sirupsen/logrus"
)

// Debug 開発中のデバッグ情報
func Debug(msg string) {
	log.WithFields(log.Fields{}).Debug(msg)
}

// Error 直ちに対処する必要のない実行時エラー
func Error(err error) {
	log.WithFields(log.Fields{}).Error(err)
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
func Warn(msg string) {
	log.WithFields(log.Fields{}).Warn(msg)
}
