package main

import(
	"log"
	"github.com/SaKu2110/sign-server/pkg/server"
	"github.com/SaKu2110/sign-server/pkg/controller"
)

func main() {
	ctrl := controller.Init()
	if err := server.Router(ctrl).Run(); err != nil {
		log.Fatal(err)
	}
}
