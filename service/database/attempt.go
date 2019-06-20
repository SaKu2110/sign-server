package database

import(
	"github.com/jinzhu/gorm"
	"context"
)

func attempt (sql *DB,ctx context.Context) {
	sql.DBMS	= "mysql"
	sql.USER	= "keeper"
	sql.PASS	= "admin_keeper"
	sql.PROTOCOL	= "tcp(127.0.0.1:3306)"
	sql.DBNAME	= "sign"
	sql.CONNECT	= sql.USER+":"+sql.PASS+"@"+sql.PROTOCOL+"/"+sql.DBNAME

	sql.DB, sql.ERROR = gorm.Open(sql.DBMS, sql.CONNECT)
}
