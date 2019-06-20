package controller

import(
	"github.com/gin-gonic/gin"
)

func (_ *CCH) PingHandler (g *gin.Context) {
	g.JSON(200, gin.H{"message":"pong"})
}
