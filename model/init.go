package model

import (
	"fmt"
	"github.com/go-ego/riot"
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
		//Roit: GetRoit(),
		Mysql: GetMysql(),
		Redis: GetRedis(),
	}
}
func (db *DataBase) Close() {

	DB.Mysql.Close()
	DB.Redis.Close()
	//DB.Roit.Close()

}

//func GetRoit() *riot.Engine {
//	// 初始化
//	searcher := riot.Engine{}
//	searcher.Init(types.EngineOpts{
//		Using:             3,
//		GseDict: "zh",
//		IndexerOpts: &types.IndexerOpts{
//			IndexType: types.LocsIndex,
//		},
//		//NotUseGse: true,
//		// GseDict: "your gopath"+"/src/github.com/go-ego/riot/data/dict/dictionary.txt",
//	})
//
//	/**
//		1. 循环读取文章,
//		2. 存入
//		3. 获取时进行裁剪
//	 */
//	file.ShowFileList(viper.GetString("passage.path"))
//
//	//text := "《复仇者联盟3：复仇无限战争》"
//	//text1 := "在IMAX影院放映时"
//	//text2 := "全片以上下扩展至IMAX 1.9：1的战争宽高比来呈现"
//
//	for index, value := range file.FileList {
//		searcher.Index(strconv.Itoa(index), types.DocData{Content: value.Content})
//	}
//
//	// 将文档加入索引，docId 从1开始
//	//searcher.Index("1", types.DocData{Content: text})
//	//searcher.Index("2", types.DocData{Content: text1}, false)
//	//searcher.Index("3", types.DocData{Content: text2}, true)
//
//	// 等待索引刷新完毕
//	searcher.Flush()
//	// engine.FlushIndex()
//	return &searcher
//}


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
