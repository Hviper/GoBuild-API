package db

import (
	"awesomeProject/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (Product) TableName() string {
	return "product"
}

type Product struct {
	Product_id            int `gorm:"primaryKey"`
	Product_name          string
	Category_id           int
	Product_title         string
	Product_intro         string
	Product_picture       string
	Product_price         float64
	Product_selling_price float64
	Product_num           int
	Product_sales         int
}

var Products []Product
var DbConnector *gorm.DB

func init() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:42128340HDC@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DbConnector = db
}

func GetDBConnector() *gorm.DB {
	return DbConnector
}

//sp_manager
func Login(users *[]common.Users) *[]common.Users {
	DbConnector.Table("sp_manager").Find(users)
	return users
}

//查询最后一个数据的信息，并返回至datastruct中
func QueryLast(tableName string, dataStruct interface{}) (result *interface{}) {
	DbConnector.Table(tableName).Last(dataStruct)
	return &dataStruct
}

//post  添加用户
func AddUser(val interface{}) {
	DbConnector.Table("db_table").Create(val)
}

//sp_user
func UsersList(userlist *[]common.UserList) *[]common.UserList {
	DbConnector.Table("sp_user").Find(userlist)
	return userlist
}
func Query() []Product {
	DbConnector.Find(&Products)
	return Products
}
