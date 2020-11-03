package controller

import(
	"fmt"
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/SaKu2110/sign-server/pkg/view"
	"github.com/SaKu2110/sign-server/pkg/model/service"
)

func (ctrl *Controller) SignInHandler (cxt *gin.Context) {
	var id, password string
	if id = cxt.GetHeader("UserId"); id == "" {
		errs := view.NewErrResponse(
			400,
			errors.New("UserId value is empty."),
		)
		cxt.JSON(http.StatusOK, view.NewSignResponse(nil, errs))
		return
	}
	if password = cxt.GetHeader("Password"); password == "" {
		errs := view.NewErrResponse(
			400,
			errors.New("Password value is empty."),
		)
		cxt.JSON(http.StatusOK, view.NewSignResponse(nil, errs))
		return
	}
	users, err := ctrl.DB.GetUserInfo(id)
	if err != nil {
		errs := view.NewErrResponse(500, err)
		cxt.JSON(http.StatusOK, view.NewSignResponse(nil, errs))
		return
	}
	if len(users) < 1 {
		errs := view.NewErrResponse(
			400,
			errors.New(fmt.Sprintf("id(%s) is not exist.", id)),
		)
		cxt.JSON(http.StatusOK, view.NewSignResponse(nil, errs))
		return
	}
	if users[0].Password != service.CreateHashWithPassord(password) {
		errs := view.NewErrResponse(
			400,
			errors.New("Incorrect password"),
		)
		cxt.JSON(http.StatusOK, view.NewSignResponse(nil, errs))
		return
	}

	// create token //
	token, err := service.CreateToken(id)
	if err != nil {
		errs := view.NewErrResponse(500, err)
		cxt.JSON(http.StatusOK, view.NewSignResponse(nil, errs))
	}
	cxt.JSON(http.StatusOK, view.NewSignResponse(&token, nil))
}
