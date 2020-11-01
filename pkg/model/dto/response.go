package dto

type Error struct {
	Code		int		`json:"code"`
	Message		string	`json:"message"`
	Description	string	`json:"description"`
}

type SignResponse struct {
	Token	*string	`json:"token"`
	Err		*Error	`json:"error"`
}
