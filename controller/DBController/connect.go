package DBController

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/SaKu2110/sign_server/model/DBModel"
)

func (dbc *DBCnt)Connect () {
	var err error
	var db *DBModel.DB
	db = &DBModel.DB{}

	// ここは後で環境変数から取ってくるようにする
	db.DBMS     = "mysql"
        db.USER     = "keeper"
        db.PASS     = "admin_keeper"
	db.PROTOCOL = "tcp(127.0.0.1:3306)"
        db.DBNAME   = "sign"
        db.CONNECT = db.USER+":"+db.PASS+"@"+db.PROTOCOL+"/"+db.DBNAME

	db.DB, err = gorm.Open(db.DBMS, db.CONNECT)
	if err != nil {
		panic(err.Error())
	}
	dbc.DB = db.DB
}
