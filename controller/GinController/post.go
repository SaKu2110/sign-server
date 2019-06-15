package GinController

import(
	"github.com/gin-gonic/gin"
)

func (gctl *GinCnt) Post (path string, handler func(*gin.Context)) {
	gctl.Gin.POST(path, handler)
}
