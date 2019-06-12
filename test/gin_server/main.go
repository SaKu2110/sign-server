package main

import "github.com/gin-gonic/gin"

type Key struct {
	Id		string `json:"id"`
	Password	string `json:"password"`
}

func GinServer() *gin.Engine {
	server := gin.Default()
	server.GET("/ping", PingResponse)
	server.POST("/signin", SigninResponse)
	server.POST("/signup", SignupResponse)
	server.Run()
	return server;
}

// /pingでアクセス
func PingResponse (g *gin.Context) {
	g.JSON(200, gin.H{ "message": "pong" })
}


// userとpassが合致した時のみトークンを発行
// ERORR: 400 リクエスト不正
func SigninResponse (g *gin.Context) {
	var key Key
	g.BindJSON(&key)
	g.JSON(200, gin.H{"access_token": "hoge"})
}

// 同一のuser名がない場合: 201
// 同一のuser名がある場合: 412
// token生成条件: 同一のDB上に存在しない場合
func SignupResponse (g *gin.Context) {
	var key Key
	g.BindJSON(&key)
	g.JSON(201, gin.H{"access_token": key.Id})
}



func main() {
    GinServer()
}
