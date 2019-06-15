package DBController

import(
	_ "github.com/jinzhu/gorm"
	"log"
)

func (dbc *DBCnt) Close () {
	dbc.DB.Close()
	log.Printf("MySQL Connect Closed!")
}
