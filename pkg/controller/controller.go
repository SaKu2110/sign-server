package controller

import(
	"github.com/SaKu2110/sign-server/pkg/model/dao"
)

type Controller struct {
	DB	*dao.DB
}

func Init(db *dao.DB) Controller {
	return Controller{DB: db}
}
