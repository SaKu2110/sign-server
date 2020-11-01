package dto

type UserInfo struct {
	ID			string	`sql:"id"`
	Password	[]byte	`sql:"password"`
}

func (u *UserInfo) CastToString() string {
	return string(u.Password)
}
