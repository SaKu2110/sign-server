package database

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	DB		*gorm.DB
	DBMS		string
	USER		string
	PASS		string
	PROTOCOL	string
	DBNAME		string
	CONNECT		string
	MSG		string
	ERROR		error
}
