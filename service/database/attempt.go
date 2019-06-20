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
	sql.DBNAME	= os.Getenv("DB_NAME")
	sql.CONNECT	= sql.USER+":"+sql.PASS+"@"+sql.PROTOCOL+"/"+sql.DBNAME

	sql.DB, sql.ERROR = gorm.Open(sql.DBMS, sql.CONNECT)
}
