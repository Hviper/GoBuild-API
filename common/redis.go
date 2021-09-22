package common

import (
	"github.com/go-redis/redis/v8"
	"time"
)

//var ctx = context.Background()
var rdb = *InitRedis()
func GetRDB() redis.Client{
	return rdb
}
func InitRedis() *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,

	})

	return rdb

	//val, err := rdb.Get(ctx, "key").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key", val)
	//
	//val2, err := rdb.Get(ctx, "key2").Result()
	//if err == redis.Nil {
	//	fmt.Println("key2 does not exist")
	//} else if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("key2", val2)
	//}

}