package controller

import (
	"net/http"

	"github.com/SaKu2110/sign-server/pkg/model/dao"
	"github.com/SaKu2110/sign-server/pkg/model/service/crypto"
	"github.com/SaKu2110/sign-server/pkg/model/service/jwt"
	"github.com/SaKu2110/sign-server/pkg/view"
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
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	if password = c.GetHeader("Password"); password == "" {
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	users, err := r.DB.GetUserInfo(id)
	if err != nil {
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	if len(users) < 1 {
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	if users[0].Password != crypto.Hash(password) {
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	token, err := jwt.Token(id)
	if err != nil {
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	c.JSON(http.StatusOK, view.NewAuthReponse(&token, nil))
}

func (r *authResolver) SignUpHandler(c *gin.Context) {
	var id, password string
	if id = c.GetHeader("UserId"); id == "" {
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	if password = c.GetHeader("Password"); password == "" {
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	users, err := r.DB.GetUserInfo(id)
	if err != nil {
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	if len(users) > 0 {
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	if err := r.DB.InsertUserInfo(
		id,
		crypto.Hash(password),
	); err != nil {
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}

	// create token //
	token, err := jwt.Token(id)
	if err != nil {
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	c.JSON(http.StatusOK, view.NewAuthReponse(&token, nil))
}
