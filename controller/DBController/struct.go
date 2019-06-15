package DBController

import(
	"github.com/jinzhu/gorm"
)

type DBCnt struct{
	DB	*gorm.DB
}
