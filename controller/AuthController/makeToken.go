package AuthController

import(
	"os"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

func (actl *AuthCnt) MakeToken (id string, action string) {
        token := jwt.New(jwt.SigningMethodHS256)

        claims := token.Claims.(jwt.MapClaims)
        claims["admin"] = true
        claims["iss"] = action
        claims["name"] = id
        claims["iat"] = time.Now()

        actl.Token, _ = token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
}
