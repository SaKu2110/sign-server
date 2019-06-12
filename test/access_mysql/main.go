package main

import (
        "encoding/json"
        "io/ioutil"
        "log"
        "github.com/jinzhu/gorm"
        _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Login struct {
        DBMS string
        USER string
        PASS string
        PROTOCOL string
        DBNAME string
}

func Connect() *gorm.DB {
        bytes, err := ioutil.ReadFile("./env/mysql/login.json")
        // Jsonファイル読み込みエラー
        if err != nil {
                log.Fatal(err)
        }
        var data Login
        // 構造体にコピー処理時のエラー
        if err := json.Unmarshal(bytes, &data); err != nil {
                log.Fatal(err)
        }

        // ログインデータの格納
        DBMS     := data.DBMS
        USER     := data.USER
        PASS     := data.PASS
        PROTOCOL := data.PROTOCOL
        DBNAME   := data.DBNAME

        CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
        db,err := gorm.Open(DBMS, CONNECT)
        // RDBとの接続を確立時のエラー
        if err != nil {
                panic(err.Error())
        }
        return db
}

func main(){
	db := Connect()
	defer db.Close()
}
