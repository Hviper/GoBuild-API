package controller

import (
	"awesomeProject/common"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	ID       int
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
//查找用户列表
func UserList(c *gin.Context) {
	var res []LoginForm
	db := common.GetDB()
	db.Table("db_table").Find(&res)

	c.JSON(200, gin.H{"data": res, "meta": map[string]interface{}{
		"msg":  "登录成功",
		"code": 200,
	}})

}
//用于查找最后一个用户的信息
func selectLastUser() LoginForm{
	var user LoginForm
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
		c.JSON(200,gin.H{
			"data":"",
			"meta":map[string]interface{}{
				"msg":"更新失败",
				"code":401,
			},
		})
		return
	}
	db := common.GetDB()
	db.Table("db_table").Model(&target).Update("password", newPassword)
	target.Password=newPassword
	c.JSON(200,gin.H{
		"data":target,
		"meta":map[string]interface{}{
			"msg":"更新成功",
			"code":200,
		},
	})
}
//顺序是username/password
func findUserByFiledName(field ...string) LoginForm{
	var res []LoginForm
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
	return LoginForm{}
}

//删除user ---> 指定用户名
func DelUser(c *gin.Context){
	var form LoginForm
	if c.ShouldBind(&form) == nil {
		db := common.GetDB()
		target := findUserByFiledName(form.Username,form.Password)
		if target.Username=="" || target.Password==""{
			c.JSON(200,gin.H{
				"data":nil,
				"meta":map[string]interface{}{
					"msg":"删除失败",
					"code":401,
				},
			})
			return
		}
		db.Table("db_table").Delete(&target)
		c.JSON(200,gin.H{
			"data":target,
			"meta":map[string]interface{}{
				"msg":"删除成功",
				"code":201,
			},
		})
		return
	}
	c.JSON(200,gin.H{
		"data":nil,
		"meta":map[string]interface{}{
			"msg":"删除失败",
			"code":401,
		},
	})
}
//添加//注册user
func AddUser(c *gin.Context){
	var user LoginForm
	if c.ShouldBind(&user)== nil {
		//最后一个id
		lastUser := selectLastUser()
		user.ID = lastUser.ID+1
		db := common.GetDB()
		db.Table("db_table").Create(&user)
		c.JSON(200,gin.H{
			"data":user,
			"meta":map[string]interface{}{
				"msg":"注册成功",
				"code":200,
			},
		})
		return
	}
	c.JSON(200,gin.H{
		"data":nil,
		"meta":map[string]interface{}{
			"msg":"注册失败",
			"code":401,
		},
	})
}

//登录
func UserLogin(c *gin.Context) {
	var form LoginForm
	var res []LoginForm
	if c.ShouldBind(&form) == nil {
		db := common.GetDB()
		db.Table("db_table").Find(&res)
		for _, v := range res {
			if v.Username == form.Username && v.Password == form.Password {
				c.JSON(200, gin.H{"data": "null", "meta": map[string]interface{}{
					"msg":  "登录成功",
					"code": 200,
				}})
				return
			}
		}
		c.JSON(401, gin.H{"data": "null", "meta": map[string]interface{}{
			"msg":  "登录失败",
			"code": 404,
		}})
	}
}
