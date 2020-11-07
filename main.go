package main

import (
	"os"

	"github.com/SaKu2110/sign-server/pkg/model/service/log"

	"github.com/SaKu2110/sign-server/pkg/controller"
	"github.com/SaKu2110/sign-server/pkg/model/dao/database"
	"github.com/SaKu2110/sign-server/pkg/server"
)

func main() {
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer db.Close()
	ctrl := controller.New(db)
	if err := server.Router(ctrl).Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
