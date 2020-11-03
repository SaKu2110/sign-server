package controller

import(
	"fmt"
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/SaKu2110/sign-server/pkg/view"
	"github.com/SaKu2110/sign-server/pkg/model/service"
)

func (ctrl *Controller) SignUpHandler (cxt *gin.Context) {
	var id, password string
	if id = cxt.GetHeader("UserId"); id == "" {
		service.ShowLogInfo(
			"Request Header `UserId` value is empty.`",
		)
		errs := view.NewErrResponse(
			400,
			errors.New("UserId value is empty."),
		)
		cxt.JSON(http.StatusOK, view.NewSignResponse(nil, errs))
		return
	}
	if password = cxt.GetHeader("Password"); password == "" {
		service.ShowLogInfo(
			"Request Header `Password` value is empty.`",
		)
		errs := view.NewErrResponse(
			400,
			errors.New("Password value is empty."),
		)
		cxt.JSON(http.StatusOK, view.NewSignResponse(nil, errs))
		return
	}
	users, err := ctrl.DB.GetUserInfo(id)
	if err != nil {
		service.ShowLogWarn(err)
		errs := view.NewErrResponse(500, err)
		cxt.JSON(http.StatusOK, view.NewSignResponse(nil, errs))
		return
	}
	if len(users) > 0 {
		service.ShowLogInfo(
			fmt.Sprintf("id(%s) is already exist.", id),
		)
		errs := view.NewErrResponse(
			400,
			errors.New(fmt.Sprintf("id(%s) is already exist.", id)),
		)
		cxt.JSON(http.StatusOK, view.NewSignResponse(nil, errs))
		return
	}
	if err := ctrl.DB.InsertUserInfo(
			id,
			service.CreateHashWithPassord(password),
		); err != nil {
			service.ShowLogWarn(err)
			errs := view.NewErrResponse(500, err)
			cxt.JSON(http.StatusOK, view.NewSignResponse(nil, errs))
			return
	}

	// create token //
	token, err := service.CreateToken(id)
	if err != nil {
		service.ShowLogWarn(err)
		errs := view.NewErrResponse(500, err)
		cxt.JSON(http.StatusOK, view.NewSignResponse(nil, errs))
		return
	}
	cxt.JSON(http.StatusOK, view.NewSignResponse(&token, nil))
}
