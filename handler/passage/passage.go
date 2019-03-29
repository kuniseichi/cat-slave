package passage

import (
	passageDao "cat-slave/model/passage"
	result "cat-slave/utils/http"
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
func List(g *gin.Context) {
	passages, err := passageDao.List()
	if err != nil {
		result.UError(g, "查询失败")
		return
	}
	result.Success(g, map[string]interface{}{
		"passages": passages,
	})

}

// 文章列表
func ListTest(g *gin.Context) {

	result.Success(g, map[string]interface{}{
		"passages": "1231241515",
	})

}

// 全文索引
