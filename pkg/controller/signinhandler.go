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
		cxt.JSON(
			http.StatusBadRequest,
			view.MakeSignResponse(
				http.StatusBadRequest, nil,
				errors.New("UserId value is empty."),
			),
		)
		return
	}
	if password = cxt.GetHeader("Password"); password == "" {
		cxt.JSON(
			http.StatusBadRequest,
			view.MakeSignResponse(
				http.StatusBadRequest, nil,
				errors.New("Password value is empty."),
			),
		)
		return
	}
	users, err := ctrl.DB.GetUserInfo(id)
	if err != nil {
		cxt.JSON(
			http.StatusInternalServerError,
			view.MakeSignResponse(
				http.StatusInternalServerError,
				nil, err,
			),
		)
		return
	}
	if len(users) < 1 {
		cxt.JSON(
			http.StatusBadRequest,
			view.MakeSignResponse(
				http.StatusBadRequest, nil,
				errors.New(fmt.Sprintf("id(%s) is not exist.", id)),
			),
		)
		return
	}
	if users[0].Password != service.CreateHashWithPassord(password) {
		cxt.JSON(
			http.StatusBadRequest,
			view.MakeSignResponse(
				http.StatusBadRequest, nil,
				errors.New("Incorrect password"),
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
				http.StatusInternalServerError,
				nil, err,
			),
		)
	}
	cxt.JSON(
		http.StatusOK,
		view.MakeSignResponse(http.StatusOK, &token, nil),
	)
}
