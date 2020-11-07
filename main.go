package main

import (
	"log"

	"github.com/SaKu2110/sign-server/pkg/controller"
	"github.com/SaKu2110/sign-server/pkg/model/dao/database"
	"github.com/SaKu2110/sign-server/pkg/server"
)

func main() {
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ctrl := controller.New(db)
	if err := server.Router(ctrl).Run(); err != nil {
		log.Fatal(err)
	}
}
