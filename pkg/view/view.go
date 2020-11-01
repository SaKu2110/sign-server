package view

import(
	"fmt"
	"github.com/SaKu2110/sign-server/pkg/model/dto"
)

const(
	STATUS_CODE_400 = "Bad Request"
	STATUS_CODE_500 = "Internal Server Error"
)

func makeErrResponse(code int, err error) *dto.Error {
	var msg string
	switch code {
	case 200:
		return nil
	case 400:
		msg = STATUS_CODE_400
	case 500:
		msg = STATUS_CODE_500
	}
	return &dto.Error{
		Code: code,
		Message: msg,
		Description: fmt.Sprintf("%v", err),
	}
}

func MakeSignResponse(code int, token *string, err error) dto.SignResponse {
	return dto.SignResponse{Token: token, Err: makeErrResponse(code, err)}
}
