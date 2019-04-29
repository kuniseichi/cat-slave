package index

import (
	indexDao "cat-slave/model/index"
	"cat-slave/pkg/http/result"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 全文索引
func TotalIndex(g *gin.Context) {

	fmt.Println(g.Param("keyword"))
	result.Success(g, map[string]interface{}{
		"passages": indexDao.TotalIndex(g.Param("keyword")),
	})

}