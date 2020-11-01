package main

import(
	"log"
	"github.com/SaKu2110/sign-server/pkg/server"
	"github.com/SaKu2110/sign-server/pkg/model/dao"
	"github.com/SaKu2110/sign-server/pkg/controller"
)

func main() {
	db, err := dao.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ctrl := controller.Init(db)
	if err := server.Router(ctrl).Run(":9000"); err != nil {
		log.Fatal(err)
	}
}
