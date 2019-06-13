package server

import(
        "github.com/gin-gonic/gin"
        "../access_mysql"
        "../auth"
)

type sign struct {
        Id              string `json:"id"`
        Password        string `json:"password"`
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
        db := access_mysql.Connect()
        defer db.Close()

        // Json形式のデータを構造体の中に格納
        var key sign
        var user sign
        g.BindJSON(&key)

        db.Find(&user, "id=? and password=?", key.Id, key.Password)
        if user.Id == key.Id && user.Password == key.Password {
                token, _ := auth.MakeToken(key.Id)
                g.JSON(200, gin.H{"access_token": token})
        }else{
                g.JSON(400, gin.H{"message": "リクエスト不正"})
        }
}

// 同一のuser名がない場合: 201
// 同一のuser名がある場合: 412
// token生成条件: 同一のDB上に存在しない場合
func SignupResponse (g *gin.Context) {
        db := access_mysql.Connect()
        defer db.Close()

        // Json形式のデータを構造体の中に格納
        var key sign
        var user sign
        g.BindJSON(&key)

        db.Find(&user, "id=?", key.Id)
        if user.Id == key.Id {
                g.JSON(412, gin.H{"massage": "同一名のユーザーが存在します"})
        }else{
                db.Create(&key)
                token, _ := auth.MakeToken(key.Id)
                g.JSON(201, gin.H{"access_token": token})
        }
}
