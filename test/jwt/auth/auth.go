package auth

import (
    "log"
    "net/http"
    "os"
    "time"

    jwtmiddleware "github.com/auth0/go-jwt-middleware"
    jwt "github.com/dgrijalva/jwt-go"
)

// GetTokenHandler get token
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)
    claims["admin"] = true
    claims["iss"] = "keeper"
    claims["name"] = "cloud-fun"
    claims["iat"] = time.Now()

    tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
    log.Printf(tokenString)
    w.Write([]byte(tokenString))
})

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
    ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("SIGNINGKEY")), nil
    },
    SigningMethod: jwt.SigningMethodHS256,
})
