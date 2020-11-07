package controller

import (
	"net/http"

	"github.com/SaKu2110/sign-server/pkg/model/dao/database"
	"github.com/SaKu2110/sign-server/pkg/model/service/auth"
	"github.com/SaKu2110/sign-server/pkg/model/service/jwt"
	"github.com/SaKu2110/sign-server/pkg/view"
	"github.com/gin-gonic/gin"
)

type AuthResolver interface {
	SignInHandler(*gin.Context)
	SignUpHandler(*gin.Context)
}

type authResolver struct {
	UserDB database.UserRepository
}

func (c *Controller) Auth() AuthResolver {
	return &authResolver{UserDB: c.UserDB}
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
	users, err := r.UserDB.Get(id)
	if err != nil {
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	if len(users) < 1 {
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	if !auth.CheckPassword(users[0], password) {
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
	users, err := r.UserDB.Get(id)
	if err != nil {
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	if len(users) > 0 {
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}

	password = auth.Hash(password)
	if err := r.UserDB.Add(id, password); err != nil {
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
