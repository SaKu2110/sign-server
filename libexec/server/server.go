package server

import(
	"github.com/gin-gonic/gin"
	"../access_mysql"
)

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
	db := access_mysql.Connect()

	// Json形式のデータを構造体の中に格納
	var key Key
	g.BindJSON(&key)

	rows, err := db.Query("select id, password from user_db")
	defer db.Close()
	g.JSON(200, gin.H{"access_token": "hoge"})
}

// 同一のuser名がない場合: 201
// 同一のuser名がある場合: 412
// token生成条件: 同一のDB上に存在しない場合
func SignupResponse (g *gin.Context) {
	db := access_mysql.Connect()

	// Json形式のデータを構造体の中に格納
	var key Key
	g.BindJSON(&key)

	rows, err := db.Query("select id from user")
	// 取ってくるのに失敗
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id string
		err := rows.Scan(id)
		// データ移しに失敗
		if err != nil {
			log.Fatal(err)
		}else if id == key.Id {
			g.JSON(401, gin.H{"message":"すでに同じユーザー名が存在します"})
		}
	}
	defer db.Close()
	g.JSON(201, gin.H{"access_token": key.Id})
}
