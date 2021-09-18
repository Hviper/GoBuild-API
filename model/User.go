package model


type LoginForm struct {
	Model    `gorm:"embedded"`
	ID        uint `gorm:"primarykey"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
