# Authentication Server for SaKu.rb 〜永遠に続く、たった1日の恋〜 Special Summer Edition
## 概要
ict演習サーバー班　課題  
- golang: go1.12.5 linux/amd64
- mysql: 14.14 Distrib 5.7.26, for Linux (x86_64) using  EditLine wrapper
## 実装した機能
- [GET] /ping  
  - Request: なし  
  - Response: {"message": "ping"} / HttpStatusCode(200)
- [POST] /signin
  - Request: {"id": 任意の文字列, "password": 任意の文字列}
  - Response: {"access_token": 任意の文字列} / HttpStatusCode(200)  

※ access_tokenは，ログインするユーザ毎に重複がないようにすること
- [POST] /signup
  - Request: {"id": 任意の文字列, "password": 任意の文字列}
  - Response: {"access_token": 任意の文字列} / HttpStatusCode(201)  

※ access_tokenは，ログインするユーザ毎に重複がないようにすること

## 実行手順
0. `setup/lib/lib.sh`で必要なライブラリを落としてくる
1. `exec`ディレクトリ内で`go run main.go`を実行
2. `curl`コマンドで遊ぶ

## ファイル構成
- ### env: 
  コードに直接書きたくない情報をJSONにしたためています  
  - #### mysql: MySQLにアクセスするユーザーとログイン情報
- ### exec: 
  main.goが格納されているファイル。ここでプログラムを実行します。
- ### libexec: 
  main.goの起動に必要なパッケージを呼び出します。
  - #### server: ginサーバーを立てる
  - #### access_mysql: mysqlとの接続を確立します
  - #### auth: JWTでトークンを生成しています
- ### setup: 
  プログラムを実行するための諸々を用意するshellが入っています
- ### test: 
  ノリで書いたテストコードです。嘘です。テストコードもどきです。
