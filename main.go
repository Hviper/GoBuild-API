package main

import (
	"awesomeProject/common"
	"awesomeProject/middleware"
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
	//加载配置文件初始化
	InitConfig()

	//MySQL初始化
	//db :=*common.InitDB()
	common.InitDB()

	//路由初始化
	r := gin.Default()
	//设置跨域允许设置
	r.Use(middleware.CorsHandler())
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
