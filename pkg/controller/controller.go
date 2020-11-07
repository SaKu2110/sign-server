package controller

import "github.com/SaKu2110/sign-server/pkg/model/dao/database"

type Controller struct {
	UserDB database.UserRepository
}

func New(db *database.DB) Controller {
	return Controller{UserDB: db.UserDB}
}
