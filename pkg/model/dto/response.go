package dto

type Error struct {
	Code		int		`json:"code"`
	Description	string	`json:"description"`
}

type SignResponse struct {
	Token	*string	`json:"token"`
	Err		*Error	`json:"error"`
}
