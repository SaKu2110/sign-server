package view

import (
	"github.com/SaKu2110/sign-server/pkg/model/dto"
)

func NewError(code int) *dto.Error {
	msg := errorMessage(code)
	return &dto.Error{Code: code, Description: msg}
}

func errorMessage(code int) string {
	var msg string
	switch code % 1000 {
	case 401:
		msg = dto.ERR_CODE_401
	case 411:
		msg = dto.ERR_CODE_411
	case 500:
		msg = dto.ERR_CODE_500
	}
	return msg
}

// NewAuthReponse returns sign handler response
func NewAuthReponse(token *string, err *dto.Error) dto.AuthResponse {
	if err != nil {
		return dto.AuthResponse{Token: "", Err: err}
	}
	return dto.AuthResponse{Token: *token, Err: nil}
}
