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
		err := errors.New("UserId value is empty.")
		cxt.JSON(
			http.StatusBadRequest,
			view.MakeSignResponse(
				nil,
				view.MakeErrResponse(http.StatusBadRequest, err),
			),
		)
		return
	}
	if password = cxt.GetHeader("Password"); password == "" {
		err := errors.New("Password value is empty.")
		cxt.JSON(
			http.StatusBadRequest,
			view.MakeSignResponse(
				nil,
				view.MakeErrResponse(http.StatusBadRequest, err),
			),
		)
		return
	}
	users, err := ctrl.DB.GetUserInfo(id)
	if err != nil {
		cxt.JSON(
			http.StatusInternalServerError,
			view.MakeSignResponse(
				nil,
				view.MakeErrResponse(http.StatusInternalServerError, err),
			),
		)
		return
	}
	if len(users) > 0 {
		err := errors.New(fmt.Sprintf("id(%s) is already exist.", id))
		cxt.JSON(
			http.StatusBadRequest,
			view.MakeSignResponse(
				nil,
				view.MakeErrResponse(http.StatusBadRequest, err),
			),
		)
		return
	}
	if err := ctrl.DB.InsertUserInfo(
			id,
			service.CreateHashWithPassord(password),
		); err != nil {
			cxt.JSON(
				http.StatusInternalServerError,
				view.MakeSignResponse(
					nil,
					view.MakeErrResponse(http.StatusInternalServerError, err),
				),
			)
			return
	}

	// create token //
	token, err := service.CreateToken(id)
	if err != nil {
		cxt.JSON(
			http.StatusInternalServerError,
			view.MakeSignResponse(
				nil,
				view.MakeErrResponse(http.StatusInternalServerError, err),
			),
		)
	}
	cxt.JSON(http.StatusOK, view.MakeSignResponse(&token, nil))
}
