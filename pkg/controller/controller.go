package controller

import(
	"github.com/SaKu2110/sign-server/pkg/model/dao"
)

type Controller struct {
	DB	*dao.DB
}

func Init() Controller {
	return Controller{}
}
