package service

import(
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
    claims["id"] = id
    claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Minute).Unix()
	
	return token.SignedString([]byte(time.Now().String()))
}
