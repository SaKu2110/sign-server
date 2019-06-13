package auth

import (
        "os"
        "time"
        jwt "github.com/dgrijalva/jwt-go"
)

func MakeToken (id string) (string, error) {

        token := jwt.New(jwt.SigningMethodHS256)

        claims := token.Claims.(jwt.MapClaims)
        claims["admin"] = true
        claims["iss"] = "keeper"
        claims["name"] = id
        claims["iat"] = time.Now()

        return token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
}

