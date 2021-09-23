package controller

import (
	"awesomeProject/common"
	"awesomeProject/model"
	"awesomeProject/response"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

//支持分页查询
func QueryNewsList(c *gin.Context) {
	db := common.GetDB()
	newsdb := db.Table("news_list").Model(&model.News{})
	var count int64
	newsdb.Count(&count) //总行数
	//post请求中获取
	pageindex, err1 := strconv.Atoi(c.PostForm("pageIndex"))
	pagesize, err2 := strconv.Atoi(c.PostForm("pageSize"))
	if err1 != nil || err2 != nil {
		response.Fail(c, gin.H{}, "请携带data字段pageIndex和pageSize", response.BADRE_QUEST,map[string]interface{}{
			"result":"请携带data字段pageIndex和pageSize",
		})
		return
	}
	newsList := make([]model.News, 0)
	if pageindex > 0 && pagesize > 0 {
		newsdb.Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&newsList)
		response.Success(c, gin.H{
			"data":  newsList,
			"count": count,
		}, "获取成功", response.OK,map[string]interface{}{
			"result":"QueryNewsList获取成功",
		})
		return
	}
	response.Fail(c, gin.H{}, "请填写合法的pageIndex和pageSize", response.BADRE_QUEST,map[string]interface{}{
		"result":"QueryNewsList请填写合法的pageIndex和pageSize",
	})
}

//del by ID
func DelNews(c *gin.Context) {
	var target model.News
	id, e := strconv.Atoi(c.PostForm("id"))
	if e != nil {
		response.Fail(c, gin.H{}, "填写参数有误", response.BADRE_QUEST,map[string]interface{}{
			"result":"QueryNewsList填写参数有误",
		})
		return
	}
	target.ID = id
	db := common.GetDB()
	// target 的 ID 是 `c.PostForm(id)`
	db.Table("news_list").Delete(&target)
	response.Success(c, gin.H{
		"data": target,
	}, "删除成功", response.OK,map[string]interface{}{
		"result":"DelNews删除成功",
	})
}

var ctx = context.Background()

//update by id
func UpdateNews(c *gin.Context) {
	//TODO

}
func AddNews(c *gin.Context) {
	//客户端传递的data对象
	var newObj model.News

	db := common.GetDB()
	//定义最后数据库最后一个节点数据
	var lastNews model.News

	//频繁使用最后一个节点存储在redis中，先尝试获取
	rdb := common.GetRDB()
	val, err := rdb.Get(ctx, "lastNews").Result()
	if err == redis.Nil {
		//fmt.Println("key2 does not exist")，缓存中不存在，需要从数据库中获取
		db.Table("news_list").Last(&lastNews)
		////拿到以后并且更新之后才能存储在redis中，需要将结构体数据转为json字符串
		//str, err := json.Marshal(lastNews)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//rdb.Set(ctx, "lastNews", str, time.Second*300)

	} else if err != nil {
		response.ServerError(c, gin.H{}, "内部解析错误", response.ERROR,map[string]interface{}{
			"result":"AddNews内部解析错误",
		})
		panic(err)
		return
	} else {
		//缓存存在，直接从redis缓存中获取数据，并 “ 反序列化 ”
		//字符串转为结构体类型数据
		str := []byte(val)
		//【反序列化为结构体】
		err := json.Unmarshal(str, &lastNews)
		if err != nil {
			response.ServerError(c, gin.H{}, "本地redis字符串解析有误错误", response.ERROR,map[string]interface{}{
				"result":"AddNews本地redis字符串解析有误错误",
			})
		}
		return
	}
	//解析请求数据  请求参数字段：title,content,timer
	if c.ShouldBind(&newObj) == nil {
		newObj.ID = lastNews.ID + 1
		fmt.Println("---------->", newObj)
		db.Table("news_list").Omit("CreatedAt", "UpdatedAt").Create(&newObj)
		//插入成功
		response.Success(c, gin.H{
			"res": newObj,
		}, "添加成功", response.OK,map[string]interface{}{
			"result":"AddNews添加成功",
		})

		//redis中的字段 “lastNews” 最后一条数据需要更新为当前插入的最新数据，
		//需要将结构体数据转为json字符串
		str, err := json.Marshal(newObj)
		if err != nil {
			fmt.Println(err)
		}
		rdb.Set(ctx, "lastNews", str, time.Second*300)
	}else{
		response.ServerError(c, gin.H{
			"res": "失败",
		}, "添加失败", response.ERROR,map[string]interface{}{
			"result":"AddNews添加失败",
		})
	}
}
