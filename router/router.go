package router

import (
	"awesomeProject/controller"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

func RouterGroup(router *gin.Engine) *gin.Engine{
	//list := db.Query()
	user := router.Group("/api/private/v1/user")

	{
		user.POST("/login",controller.UserLogin)
		opt := user.Group("")
		opt.Use(middleware.AdminMiddleware())
		{
			opt.POST("/delUser",controller.DelUser)
			opt.POST("/updateUser",controller.UpdateUser)
			opt.POST("/addUser",controller.AddUser)
			opt.GET("/userList",controller.UserList)
		}

	}

	goods := router.Group("/api/private/v1/goods")
	{
		goods.GET("/select",controller.SelectGoods)
		goods.POST("/delGoods",controller.DelGoods)
		goods.POST("/updateGoods",controller.UpdateGoods)
		goods.POST("/addGoods",controller.AddGoods)
		goods.GET("/goods",controller.Goods)

	}

	return router
}