package models

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var RedisDB *redis.Client

func init() {

	db, err := gorm.Open("mysql", "root:yingli@tcp(10.10.1.101:3306)/gubao?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	DB = db
	DB.DB().SetConnMaxLifetime(600) //TODO::这个地方有可能产生bug
	DB.DB().SetMaxOpenConns(128)
	DB.DB().SetMaxIdleConns(16)
	//defer DB.Close()

	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "10.10.1.101:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}
