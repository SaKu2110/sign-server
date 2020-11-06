package view

import (
	"github.com/SaKu2110/sign-server/pkg/model/dto"
)

// NewAuthReponse returns sign handler response
func NewAuthReponse(token *string, code *string) dto.AuthResponse {
	if code != nil {
		return dto.AuthResponse{}
	}
	if token == nil {
		return dto.AuthResponse{}
	}
	return dto.AuthResponse{Token: *token}
}
