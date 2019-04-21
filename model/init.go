package model

import (
	"fmt"
	"github.com/go-ego/riot"
	"github.com/olivere/elastic"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type DataBase struct {
	Mysql *gorm.DB
	Redis *redis.Client
	Roit  *riot.Engine
}

var DB *DataBase

func (db *DataBase) Init() {
	DB = &DataBase{
		Roit: GetRoit(),
		//Mysql: GetMysql(),
		// Redis: GetRedis(),
	}
}
func (db *DataBase) Close() {

	//DB.Mysql.Close()
	// DB.Redis.Close()
}

func GetRoit() *riot.Engine {

	return nil
}

func openElastic() *elastic.Client {
	return nil
}

func GetMysql() *gorm.DB {
	return openMysql(viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.addr"),
		viper.GetString("mysql.name"))
}
func GetRedis() *redis.Client {
	return openRedis(viper.GetString("redis.addr"),
		viper.GetString("redis.password"))
}
func openRedis(addr, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Errorf(err, "redis connection failed")
	}
	return client
}

func openMysql(username, password, addr, name string) *gorm.DB {

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	setupMysql(db)

	return db

}

func setupMysql(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxIdleConns(0)
}
