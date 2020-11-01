package server

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/SaKu2110/sign-server/pkg/controller"
)

func Router(ctrl controller.Controller) (router *gin.Engine) {
	router = gin.Default()
	// handler作るのめんどくさかった...
	router.GET("/ping", func(cxt *gin.Context) {
        cxt.JSON(http.StatusOK, gin.H{"message": "pong"})
    })
	router.POST("/signin", ctrl.SignInHandler)
	router.POST("/signup", ctrl.SignUpHandler)
	return
}
