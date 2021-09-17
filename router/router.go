package router

import (
	"github.com/gin-gonic/gin"
	"awesomeProject/controller"
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
		goods.POST("/select")
		goods.POST("/delGoods")
		goods.POST("/updateGoods")
		goods.POST("/addGoods")
		goods.GET("/goods")

	}

	return router
}