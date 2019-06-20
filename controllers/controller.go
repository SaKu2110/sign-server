package controller

import(
	"github.com/jinzhu/gorm"
)

type CCH struct {
	DB	*gorm.DB
	ERROR	error
}
