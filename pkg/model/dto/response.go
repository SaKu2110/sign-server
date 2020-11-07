package dto

// AuthResponse is auth handler response
type AuthResponse struct {
	Token string `json:"token,omitempty"`
	Err   *Error `json:"error,omitempty"`
}
