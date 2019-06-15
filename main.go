package main

import(
	"github.com/SaKu2110/sign_server/controller/DBController"
	"github.com/SaKu2110/sign_server/controller/GinController"
	"github.com/SaKu2110/sign_server/controller/ResController"
)

var(
	DB = &DBController.DBCnt{}
	Gin = &GinController.GinCnt{}
	Res = &ResController.ResCnt{}
)

func init() {
	DB.Connect()
}

func main() {
	Gin.Init()
	Gin.Get("/ping", Res.PingHandler)
	Gin.Post("/signin")
	Gin.Post("/signup")
	Gin.Run()
	defer DB.Close()
}
