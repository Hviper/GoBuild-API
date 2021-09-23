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
		//用户级别的操作需要使用中间件来控制，token
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
		goods.GET("/goodsList",controller.Goods)

	}

	newslist := router.Group("/api/private/v1/news")
	{
		//查询可以指定分页情况
		newslist.POST("/query",controller.QueryNewsList)
		newslist.POST("/delNews",controller.DelNews)
		newslist.POST("/updateNews",controller.UpdateNews)
		newslist.POST("/addNews",controller.AddNews)

	}

	group_dynamics := router.Group("/api/private/v1/group_dynamics")
	{
		//try test the interface
		group_dynamics.GET("/query",controller.QueryGroup_Dynamics)
	}

	return router
}