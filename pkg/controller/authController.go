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
		errs := view.NewError(101411)
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, errs))
		return
	}
	if password = c.GetHeader("Password"); password == "" {
		log.Info("Password value is empty.")
		errs := view.NewError(101411)
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, errs))
		return
	}
	users, err := r.UserDB.Get(id)
	if err != nil {
		log.Error(err)
		errs := view.NewError(301500)
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, errs))
		return
	}
	if len(users) < 1 {
		log.Info(fmt.Sprintf("There was no corresponding data about ID(%s) in the `user` Database.", id))
		errs := view.NewError(101401)
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, errs))
		return
	}
	if !auth.CheckPassword(users[0], password) {
		log.Info("Password is incorrect.")
		errs := view.NewError(101401)
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, errs))
		return
	}

	// create token //
	token, err := jwt.Token(id)
	if err != nil {
		log.Error(err)
		errs := view.NewError(501500)
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, errs))
		return
	}
	c.JSON(http.StatusOK, view.NewAuthReponse(&token, nil))
}

func (r *authResolver) SignUpHandler(c *gin.Context) {
	log.Debug("Get /signup request")
	var id, password string
	if id = c.GetHeader("UserId"); id == "" {
		log.Info("UserId value is empty.")
		errs := view.NewError(102411)
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, errs))
		return
	}
	if password = c.GetHeader("Password"); password == "" {
		log.Info("Passwrod value is empty.")
		errs := view.NewError(102411)
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, errs))
		return
	}
	users, err := r.UserDB.Get(id)
	if err != nil {
		log.Error(err)
		errs := view.NewError(302500)
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, errs))
		return
	}
	if l := len(users); l > 0 {
		log.Info(fmt.Sprintf("There was %d corresponding data about ID(%s) in the `user` Database.", l, id))
		errs := view.NewError(102401)
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, errs))
		return
	}

	password = auth.Hash(password)
	if err := r.UserDB.Add(id, password); err != nil {
		log.Error(err)
		errs := view.NewError(302500)
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, errs))
		return
	}
	// create token //
	token, err := jwt.Token(id)
	if err != nil {
		log.Error(err)
		errs := view.NewError(502500)
		c.JSON(http.StatusOK, view.NewAuthReponse(nil, errs))
		return
	}
	c.JSON(http.StatusOK, view.NewAuthReponse(&token, nil))
}
