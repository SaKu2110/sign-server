package dto

// Errors is error data loader
type Errors map[string]Error

// Error is error response data format
type Error struct {
	Code int `json:"code"`
}
