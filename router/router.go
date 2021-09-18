package router

import (
	"awesomeProject/controller"
	"github.com/gin-gonic/gin"
)

func RouterGroup(router *gin.Engine) *gin.Engine{
	//list := db.Query()
	user := router.Group("/api/private/v1/user")
	{
		user.POST("/login",controller.UserLogin)
		user.POST("/delUser",controller.DelUser)
		user.POST("/updateUser",controller.UpdateUser)
		user.POST("/addUser",controller.AddUser)
		user.GET("/userList",controller.UserList)


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