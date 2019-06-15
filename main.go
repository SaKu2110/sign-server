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
	Res.DB = DB.DB
}

func main() {
	Gin.Init()
	Gin.Get("/ping", Res.PingHandler)
	Gin.Post("/signin", Res.SigninHandler)
	Gin.Post("/signup", Res.SignupHandler)
	Gin.Run()
	defer DB.Close()
}
