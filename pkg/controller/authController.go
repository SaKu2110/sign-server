package controller

import (
	"fmt"
	"net/http"

	"github.com/SaKu2110/sign-server/pkg/model/dao/database"
	"github.com/SaKu2110/sign-server/pkg/model/service/auth"
	"github.com/SaKu2110/sign-server/pkg/model/service/jwt"
	"github.com/SaKu2110/sign-server/pkg/model/service/log"
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
		log.Info("UserId value is empty.")
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	if password = c.GetHeader("Password"); password == "" {
		log.Info("Password value is empty.")
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	users, err := r.UserDB.Get(id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	if len(users) < 1 {
		log.Info(fmt.Sprintf("There was no corresponding data about ID(%s) in the `user` Database.", id))
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	if !auth.CheckPassword(users[0], password) {
		log.Info("Password is incorrect.")
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}

	// create token //
	token, err := jwt.Token(id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	c.JSON(http.StatusOK, view.NewAuthReponse(&token, nil))
}

func (r *authResolver) SignUpHandler(c *gin.Context) {
	var id, password string
	if id = c.GetHeader("UserId"); id == "" {
		log.Info("UserId value is empty.")
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	if password = c.GetHeader("Password"); password == "" {
		log.Info("Passwrod value is empty.")
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	users, err := r.UserDB.Get(id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	if l := len(users); l > 0 {
		log.Info(fmt.Sprintf("There was %d corresponding data about ID(%s) in the `user` Database.", l, id))
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}

	password = auth.Hash(password)
	if err := r.UserDB.Add(id, password); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}

	// create token //
	token, err := jwt.Token(id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, nil))
		return
	}
	c.JSON(http.StatusOK, view.NewAuthReponse(&token, nil))
}
