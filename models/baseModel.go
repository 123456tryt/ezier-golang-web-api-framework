package models

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var RedisDB *redis.Client

func init() {
	DB, _ = gorm.Open("mysql", "alialiali:xxxxx@/www.xxxx.org?charset=utf8&parseTime=True&loc=Local")
	defer DB.Close()

	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}
