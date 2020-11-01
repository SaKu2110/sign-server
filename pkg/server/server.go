package server

import(
	"github.com/gin-gonic/gin"
	"github.com/SaKu2110/sign-server/pkg/controller"
)

func Router(ctrl controller.Controller) (router *gin.Engine) {
	router = gin.Default()
	router.GET("/ping", ctrl.PingHandler)
	router.POST("/signin", ctrl.SignInHandler)
	router.POST("/signup", ctrl.SignUpHandler)
	return
}
