package GinController

func (gctl *GinCnt) Post (path string) {
	gctl.Gin.POST(path)
}
