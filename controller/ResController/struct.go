package ResController

import(
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/SaKu2110/sign_server/controller/AuthController"
)

type ResCnt struct {
	Cntxt	*gin.Context
	DB	*gorm.DB
	AuthController.AuthCnt
}
