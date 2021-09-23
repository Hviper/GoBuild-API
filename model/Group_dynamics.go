package model
type Group_dynamics struct {
	Model    `gorm:"embedded"`
	ID int  `gorm:"primaryKey" gorm:"column:ID"`
	Title   string `gorm:"column:title"`
	ImgUrl  string `gorm:"column:img_url"`
	Content string `gorm:"column:content"`
	Timer string  `gorm:"column:timer"`
}