package model

import(
	"github.com/gin-gonic/gin"
)

type sign struct {
        Id              string `json:"id"`
        Password        string `json:"password"`
	Gin		*gin.Context
}
