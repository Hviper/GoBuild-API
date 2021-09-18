package model
type Goods struct {
	Model    `gorm:"embedded"`
	Product_id int
	Product_name string
	Category_id int
	Product_title string
	Product_intro string
	Product_picture string
	Product_price float64
	Product_selling_price float64
	Product_num int
	Product_sales int
}