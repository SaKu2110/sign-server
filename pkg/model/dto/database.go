package dto

type UserInfo struct {
	ID			string	`sql:"id"`
	Password	[]byte	`sql:"password"`
}
