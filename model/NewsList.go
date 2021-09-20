package model

type News struct {
	Model    `gorm:"embedded"`
	ID int  `gorm:"primaryKey"`
	Title string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
	Timer string `form:"timer" binding:"required"`
}
