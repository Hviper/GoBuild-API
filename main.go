package main

import (
	"awesomeProject/common"
	"awesomeProject/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "gorm.io/gorm"
	"os"
)



type User struct{
	Id int
	Username string
	Password string
}
func main() {
	//加载配置文件
	InitConfig()
	//MySQL初始化
	//db :=*common.InitDB()
	common.InitDB()

	//db :=common.GetDB()
	//var list []User
	//db.Table("db_table").Find(&list)
	//fmt.Println(db,"==========>",list)


	r := gin.Default()
	r = router.RouterGroup(r)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

//配置文件
func InitConfig() {
	workDir, _ := os.Getwd()

	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
