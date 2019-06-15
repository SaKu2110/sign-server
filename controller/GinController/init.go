package GinController

import(
        "github.com/gin-gonic/gin"
)

func (gctl *GinCnt) Init () {
	gctl.Gin = gin.Default()
}
