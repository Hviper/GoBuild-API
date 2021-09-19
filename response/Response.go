package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//| *状态码* | *含义*                | *说明*                                              |
//| -------- | --------------------- | --------------------------------------------------- |
//| 200      | OK                    | 请求成功                                            |
//| 201      | CREATED               | 创建成功                                            |
//| 204      | DELETED               | 删除成功                                            |
//| 400      | BAD REQUEST           | 请求的地址不存在或者包含不支持的参数                |
//| 401      | UNAUTHORIZED          | 未授权                                              |
//| 403      | FORBIDDEN             | 被禁止访问                                          |
//| 404      | NOT FOUND             | 请求的资源不存在                                    |
//| 422      | Unprocesable entity   | [POST/PUT/PATCH] 当创建一个对象时，发生一个验证错误 |
//| 500      | INTERNAL SERVER ERROR | 内部错误                                            |

const(
	OK = 200
	CREATED = 201
	DELETED = 204
	BADRE_QUEST = 400   //数据库各种失败
	UNAUTHORIZED = 401
	FORBIDDEN = 403
	NOT_FOUND = 404
	ERROR = 500
)


//为ctx.JSON()中的response做统一处理,统一返回两个字段 【key：meta和data】
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"meta":map[string]interface{}{
		"code": code, "msg": msg,
	}, "data": data})
}

func Success(ctx *gin.Context, data gin.H, msg string,code int) {
	Response(ctx, http.StatusOK, code, data, msg)
}
func Fail(ctx *gin.Context, data gin.H, msg string,code int) {
	Response(ctx, http.StatusBadRequest, code, data, msg)
}
func ServerError(ctx *gin.Context, data gin.H, msg string,code int){
	Response(ctx, http.StatusBadRequest, code, data, msg)
}
