package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
