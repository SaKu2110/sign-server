package server

import (
	"net/http"

	"github.com/SaKu2110/sign-server/pkg/controller"
	"github.com/gin-gonic/gin"
)

func Router(ctrl controller.Controller) (router *gin.Engine) {
	router = gin.Default()
	router.GET("/ping", func(cxt *gin.Context) {
		cxt.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	router.POST("/signin", ctrl.Auth().SignInHandler)
	router.POST("/signup", ctrl.Auth().SignUpHandler)
	return
}
