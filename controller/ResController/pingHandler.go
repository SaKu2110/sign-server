package ResController

import(
	"github.com/gin-gonic/gin"
)

func (_rctl *ResCnt) PingHandler (g *gin.Context) {
	g.JSON(200, gin.H{ "message": "pong" })
}
