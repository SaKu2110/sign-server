package auth

import(
	"os"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

func (ctrl *Auth) GetToken (id string, action string) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["iss"] = action
	claims["name"] = id
	claims["iat"] = time.Now()

	ctrl.Token, ctrl.ERROR = token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
}
