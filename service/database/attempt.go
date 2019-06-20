package database

import(
	"github.com/jinzhu/gorm"
	"context"
	"os"
)

func attempt (sql *DB,ctx context.Context) {
	sql.DBMS	= "mysql"
	sql.USER	= os.Getenv("DB_USER")
	sql.PASS	= os.Getenv("DB_PASS")
	sql.PROTOCOL	= os.Getenv("DB_IP")
	sql.PORT	= os.Getenv("DB_PORT")
	sql.DBNAME	= os.Getenv("DB_NAME")
	sql.CONNECT	= sql.USER+":"+sql.PASS+"@tcp("+sql.PROTOCOL+":"+sql.PORT+")"+"/"+sql.DBNAME+"?charset=utf8&parseTime=True&loc=Local"

	sql.DB, sql.ERROR = gorm.Open(sql.DBMS, sql.CONNECT)
}
