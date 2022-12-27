package models

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine = Init()
var Rds = InitRedis()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@/chat?charset=utf8")
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}
	fmt.Println("成功连接数据库")
	return engine
}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
