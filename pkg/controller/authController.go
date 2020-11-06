package controller

import (
	"github.com/SaKu2110/sign-server/pkg/model/dao"
	"github.com/SaKu2110/sign-server/pkg/model/service/crypto"
	"github.com/SaKu2110/sign-server/pkg/model/service/jwt"
	"github.com/gin-gonic/gin"
)

type AuthResolver interface {
	SignInHandler(*gin.Context)
	SignUpHandler(*gin.Context)
}

type authResolver struct {
	DB *dao.DB
}

func (c *Controller) Auth() AuthResolver {
	return &authResolver{DB: c.DB}
}

func (r *authResolver) SignInHandler(c *gin.Context) {
	var id, password string
	if id = c.GetHeader("UserId"); id == "" {
		return
	}
	if password = c.GetHeader("Password"); password == "" {
		return
	}
	users, err := r.DB.GetUserInfo(id)
	if err != nil {
		return
	}
	if len(users) < 1 {
		return
	}
	if users[0].Password != crypto.Hash(password) {
		return
	}
	_, err = jwt.Token(id)
	if err != nil {
		return
	}
}

func (r *authResolver) SignUpHandler(c *gin.Context) {
	var id, password string
	if id = c.GetHeader("UserId"); id == "" {
		return
	}
	if password = c.GetHeader("Password"); password == "" {
		return
	}
	users, err := r.DB.GetUserInfo(id)
	if err != nil {
		return
	}
	if len(users) > 0 {
		return
	}
	if err := r.DB.InsertUserInfo(
		id,
		crypto.Hash(password),
	); err != nil {
		return
	}

	// create token //
	_, err = jwt.Token(id)
	if err != nil {
		return
	}
}
