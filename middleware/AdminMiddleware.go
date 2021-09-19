package middleware

import (
	"awesomeProject/common"
	"awesomeProject/model"
	"awesomeProject/response"
	"github.com/gin-gonic/gin"
)

/*********************************************************
** 函数功能: 用户认证中间件
** 日    期:2021/9/19
**********************************************************/
func AdminMiddleware() gin.HandlerFunc{

	return func(ctx *gin.Context) {
		//前端Header字段需要加Authorization这个字段：携带服务器发放的Token
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			response.Fail(ctx,gin.H{
				"data":"拒绝访问",
			},"Token值为空",400)
			//抛弃请求
			ctx.Abort()
			return
		}
		//Token中解析树Claims这个用户model数据
		Claims,err := common.ParseToken(tokenString)
		if err == nil{
			//ctx.JSON(200,gin.H{
			//	"Claims":Claims,
			//})
			var user model.LoginForm
			common.GetDB().Table("db_table").Where("id = ?", Claims.ID).First(&user)
			//数据库中存在这条数据,数据拦截放行
			if user.ID!=0{
				ctx.Next()
				return
			}
		}

		response.Fail(ctx,gin.H{

			"data":"拒绝访问",
		},"token有误，请携带正确的头字段信息",400)
		//抛弃请求
		ctx.Abort()
	}
}




