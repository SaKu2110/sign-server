# sign-server

<p align="left">
  <a href="https://github.com/actions/setup-node/actions?query=workflow%3Abuild-test"><img alt="build-test status" src="https://github.com/SaKu2110/sign-server/workflows/Go/badge.svg"></a>
</p>

## 概要
ict演習サーバー班　課題  
1年の時を経て書き直したやつ
- golang: 1.15.3 darwin/amd64

## 実装した機能
- [GET] /ping  
  - Request: なし  
  - Response: {"message": "ping"} / HttpStatusCode(200)
- [POST] /signin
  - Request: Header: UserId `string`, Passwrod `string`
  - Response: {"token": 任意の文字列} / HttpStatusCode(200)  

- [POST] /signup
  - Request: Header: UserId `string`, Passwrod `string`
  - Response: {"token": 任意の文字列} / HttpStatusCode(201)  

## 実行手順
`make run`でapiを起動する


