package ResController

import(
	"github.com/gin-gonic/gin"
)

func (rctl *ResCnt) PingHandler (_g *gin.Context) {
	rctl.Cntxt.JSON(200, gin.H{ "message": "pong" })
}
