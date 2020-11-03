package view

import(
	"fmt"
	"github.com/SaKu2110/sign-server/pkg/model/dto"
)

func NewErrResponse(code int, err error) *dto.Error {
	return &dto.Error{
		Code: code,
		Description: fmt.Sprintf("%v", err),
	}
}

func NewSignResponse(token *string, err *dto.Error) dto.SignResponse {
	return dto.SignResponse{Token: token, Err: err}
}
