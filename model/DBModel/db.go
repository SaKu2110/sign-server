package DBModel

import(
	"github.com/jinzhu/gorm"
)

type DB struct {
	DB		*gorm.DB
	DBMS		string
        USER		string
        PASS		string
        PROTOCOL	string
        DBNAME		string
	CONNECT		string
}
