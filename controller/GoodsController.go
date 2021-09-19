package controller

import (
	"awesomeProject/common"
	"awesomeProject/model"
	"awesomeProject/response"
	"github.com/gin-gonic/gin"
)


//查找全部
func SelectGoods(c *gin.Context) {
	var list []model.Goods

	db := common.GetDB()

	db.Table("product").Find(&list)
	response.Success(c,gin.H{"data": list},"获取成功",response.OK)

}

//删除
func DelGoods(c *gin.Context) {

}
//更新
func UpdateGoods(c *gin.Context) {

}
//添加
func AddGoods(c *gin.Context) {

}
//返回全部数据给它
func Goods(c *gin.Context) {
	var list []model.Goods

	db := common.GetDB()

	db.Table("product").Find(&list)
	response.Success(c,gin.H{"data": list},"获取成功",response.OK)

}