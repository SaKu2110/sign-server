package GinController

import(
	"github.com/gin-gonic/gin"
)

func (gctl *GinCnt) Get (path string, handler func(*gin.Context)) {
	gctl.Gin.GET(path, handler)
}
