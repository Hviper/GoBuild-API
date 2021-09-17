package main

import (
	"awesomeProject/common"
	"awesomeProject/db"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)



const (
	SecretKey = "welcome to hdc's gtihub"
)




func getToken() common.Token {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	tokenString, _ := token.SignedString([]byte(SecretKey))
	val := common.Token{tokenString}
	return val
}
func ValidateTokenMiddleware() {
	//TODO
}

func RouterGroup() {
	router := gin.Default()
	list := db.Query()
	gp := router.Group("/api/private/v1/")
	{
		gp.POST("/login", func(c *gin.Context) {
			var loginForm common.Users
			err := c.ShouldBind(&loginForm)
			if err != nil {
				return
			}
			var users []common.Users
			db.Login(&users)
			fmt.Println(loginForm)
			for _, v := range users {
				if v.Username == loginForm.Username && v.Password == loginForm.Password {
					c.JSON(200, gin.H{
						"meta": map[string]interface{}{
							"code": 200,
							"msg":  "成功",
						},

						"data": map[string]interface{}{
							"token": getToken(),
						},
					})
					return
				}
			}
			c.JSON(400, gin.H{
				"meta": map[string]interface{}{
					"code": 400,
					"msg":  "失败",
				},

				"data": nil,
			})
		})

		gp.POST("/users", func(c *gin.Context) {
			type Form struct {
				ID       int
				Username string `form:"username" binding:"required"`
				Password string `form:"password" binding:"required"`
			}
			//用于存储当前需要存储的节点信息
			var form Form
			//用于存储最后一个数据的信息
			var formLast Form
			if c.ShouldBind(&form) == nil {

				res := *db.QueryLast("db_table", &formLast)
				//QueryLast传递进去的是指针类型，返回出来进行强制类型转换的时候也是对应为指针类型
				t, ok := res.(*Form)
				if ok {
					form.ID = t.ID + 1
				}
				db.AddUser(&form)
				c.JSON(200, map[string]interface{}{
					"data": map[string]interface{}{
						"id":       form.ID,
						"username": form.Username,
						"password": form.Password,
					},
					"meta": map[string]interface{}{
						"msg":    "用户创建成功",
						"status": 201,
					},
				})
				return
			}
			c.JSON(401, map[string]interface{}{
				"data": nil,
				"meta": map[string]interface{}{
					"msg":    "用户创建失败",
					"status": 401,
				},
			})
		})

		gp.GET("/users", func(c *gin.Context) {
			var userlist []common.UserList
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "成功",
				"data": db.UsersList(&userlist),
			})
		})

		gp.GET("/api", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "成功",
				"data": list,
			})
		})

	}

	err := router.Run(":8888")
	if err != nil {
		return 
	}
}

func main() {

	RouterGroup()

}
