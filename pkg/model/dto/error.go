package dto

const (
	ERR_CODE_401 = "Incorrect POST data"
	ERR_CODE_411 = "Header value is not defined"
	ERR_CODE_500 = "Internal Server Error"
)

// Error is error response data format
type Error struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}
