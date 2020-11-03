package service

import(
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

func CreateJWTToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
    claims["id"] = id
    claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Minute).Unix()
	
	// 秘密鍵作るの面倒かったから許して...
	return token.SignedString([]byte(time.Now().String()))
}
