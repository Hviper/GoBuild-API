package controller

import (
	"awesomeProject/common"
	"awesomeProject/model"
	"awesomeProject/response"
	"awesomeProject/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


//查找用户列表
func UserList(c *gin.Context) {
	var res []model.LoginForm
	db := common.GetDB()
	db.Table("db_table").Find(&res)
	response.Success(c,gin.H{"data": res},"登录成功",response.OK,map[string]interface{}{
		"result":"UserList登录成功",
	})


}
//用于查找最后一个用户的信息
func selectLastUser() model.LoginForm{
	var user model.LoginForm
	db := common.GetDB()
	db.Table("db_table").Last(&user)
	return user
}



//更新user数据
func UpdateUser(c *gin.Context) {
	username :=c.PostForm("username")
	oldPassword :=c.PostForm("oldPassword")
	newPassword :=c.PostForm("newPassword")
	target := findUserByFiledName(username,oldPassword)
	if target.Username=="" || target.Password==""{
		response.Fail(c,gin.H{"data": ""},"更新失败",response.BADRE_QUEST,map[string]interface{}{
			"result":"UpdateUser更新失败",
		})
		return
	}
	db := common.GetDB()
	db.Table("db_table").Model(&target).Update("password", newPassword)
	target.Password=newPassword
	response.Success(c,gin.H{"data": target},"更新成功",response.OK,map[string]interface{}{
		"result":"UpdateUser更新成功",
	})

}
//顺序是username/password
func findUserByFiledName(field ...string) model.LoginForm{
	var res []model.LoginForm
	db := common.GetDB()
	db.Table("db_table").Find(&res)
	for _,v := range res{
		if len(field)==1{
			if v.Username==field[0]{
				return v
			}
		}else if len(field) ==2{
			if v.Username==field[0] && v.Password==field[1]{
				return v
			}
		}
	}
	//空值 --> new一个空对象
	return model.LoginForm{}
}

//删除user ---> 指定用户名
func DelUser(c *gin.Context){
	var form model.LoginForm
	if c.ShouldBind(&form) == nil {
		db := common.GetDB()
		target := findUserByFiledName(form.Username,form.Password)
		if target.Username=="" || target.Password==""{
			response.Fail(c,gin.H{"data": nil},"删除失败",response.BADRE_QUEST,map[string]interface{}{
				"result":"DelUser删除失败",
			})

			return
		}
		db.Table("db_table").Delete(&target)

		response.Success(c,gin.H{"data": target},"删除成功",response.DELETED,map[string]interface{}{
			"result":"DelUser删除成功",
		})

		return
	}
	response.Fail(c,gin.H{"data": nil},"删除失败",response.BADRE_QUEST,map[string]interface{}{
		"result":"DelUser删除失败",
	})

}
//添加//注册user
func AddUser(c *gin.Context){
	var user model.LoginForm
	if c.ShouldBind(&user)== nil {
		//最后一个id
		lastUser := selectLastUser()
		user.ID = lastUser.ID+1
		db := common.GetDB()
		db.Table("db_table").Omit("CreatedAt", "UpdatedAt").Create(&user)
		response.Success(c,gin.H{"data": user},"注册成功",response.OK,map[string]interface{}{
			"result":"AddUser注册成功",
		})

		return
	}
	response.Fail(c,gin.H{"data": nil},"注册失败",response.BADRE_QUEST,map[string]interface{}{
		"result":"AddUser注册失败",
	})

}

//登录
func UserLogin(c *gin.Context) {
	var form model.LoginForm
	var res []model.LoginForm
	if c.ShouldBind(&form) == nil {
		db := common.GetDB()
		db.Table("db_table").Find(&res)
		for _, v := range res {
			if v.Username == form.Username && v.Password == form.Password {
				token, err := common.GenerateToken(v.ID)
				if err != nil {
					//写入日志
					utils.Logger().WithFields(logrus.Fields{
						"err": err,
					}).Info("A walrus appears")
					response.ServerError(c,nil,"系统异常",response.ERROR,map[string]interface{}{
						"result":"UserLogin系统异常",
					})
					return
				}
				response.Success(c,gin.H{"data": v,"token":token},"登录成功",response.OK,map[string]interface{}{
					"result":"UserLogin登录成功",
				})

				return
			}
		}
		response.Fail(c,gin.H{"data": nil},"登录失败",response.BADRE_QUEST,map[string]interface{}{
			"result":"UserLogin登录失败",
		})

	}
	response.ServerError(c,gin.H{"data": nil},"登录失败",response.ERROR,map[string]interface{}{
		"result":"UserLogin登录失败",
	})
}
