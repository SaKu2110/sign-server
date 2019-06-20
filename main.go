package main

import(
	"github.com/SaKu2110/sign-server/service/database"
	"github.com/SaKu2110/sign-server/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

var(
	SQL = &database.DB{}
	CCH = &controller.CCH{}
)

func init() {
	SQL.Connect()
	// タイムアウト処理
	if SQL.ERROR != nil {
		log.Printf(SQL.MSG)
		panic(SQL.ERROR.Error())
	}
	// ハンドラ内でDBの操作ができるようにする
	CCH.DB = SQL.DB
}

func main() {
	// Ginサーバの起動
	gin := gin.Default()
	gin.GET("/ping", CCH.PingHandler)
	gin.POST("/signin", CCH.SignInHandler)
	gin.POST("/signup", CCH.SignUpHandler)
	if CCH.ERROR != nil {
		log.Fatal(CCH.ERROR)
	}
	gin.Run()
}
