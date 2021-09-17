package common

type Users struct {
	Username string `gorm:"column:mg_name" form:"username" binding:"required"`
	Password string `gorm:"column:mg_pwd" form:"password" binding:"required"`
}

type UserList struct {
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

type Token struct {
	Token string `json:"token"`
}