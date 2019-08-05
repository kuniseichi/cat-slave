package passage

import (
	"cat-slave/model"
	"cat-slave/service/passage"
	"encoding/json"
	// passageDao "cat-slave/model/passage"
	passageDao "cat-slave/model/passage"
	"cat-slave/pkg/http/result"
	"github.com/lexkong/log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 根据id获取文章
func Get(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		result.UError(g, "参数转换失败")
		return
	}
	p, err := passageDao.Get(id)
	if err != nil {
		// log.Info()
		result.UError(g, "db获取失败")
		return
	}
	result.Success(g, map[string]interface{}{
		"passage": p,
	})
}

// 文章列表
func BrefList(g *gin.Context) {
	passages, err := passageDao.List()
	if err != nil {
		result.UError(g, "查询异常")
		return
	}
	result.Success(g, map[string]interface{}{
		"passages": passages,
	})
}

// tranction test1
func tranone(g *gin.Context) {
	// 参数校验

	// 数据库数据校验 (避免重复等)

	// 开始事务修改

	// 返回结果
}

// 文章列表
func ListTest(g *gin.Context) {
	passage.PassageService(1)
	result.Success(g, map[string]interface{}{
		"passages": "1231241515",
	})
}

func Index(g *gin.Context)  {
	keyword := g.Param("keyword")
	var value interface{}
	if model.DB.Redis.Exists("index_" + keyword).Val() != 0 {
		v := model.DB.Redis.Get("index_" + keyword).Val()
		json.Unmarshal([]byte(v),&value)
	} else {
		value2, err := passageDao.Index(keyword)
		if err != nil {
			log.Info("获取索引失败")
		}
		str, _ := json.Marshal(value2)
		model.DB.Redis.Set("index_" + keyword, string(str), 1000*1000*1000*3600*24)
		value = value2
	}

	result.Success(g, map[string]interface{}{
		"passages": value,
	})
}


