package controller

import (
	"awesomeProject/common"
	"awesomeProject/model"
	"awesomeProject/response"
	"github.com/gin-gonic/gin"
)

func QueryGroup_Dynamics(c *gin.Context){
	db := common.GetDB()
	var Gys []model.Group_dynamics
	db.Table("group_dynamics").Find(&Gys)
	response.Success(c,gin.H{
		"data":Gys,
	},"获取成功",200,map[string]interface{}{
		"result":"QueryGroup_Dynamics 获取成功",
	})
}