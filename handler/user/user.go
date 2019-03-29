package user

import (
	userDao "cat-slave/model/user"
	"cat-slave/utils/err"
	result "cat-slave/utils/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(g *gin.Context) {
	u := userDao.GetUserList()
	result.Success(g, map[string]interface{}{
		"userList": u,
		"token":    "",
	})

}

func GetUserList2(g *gin.Context) {
	// u := userDao.GetUserList()
	result.UError(g, "")
}

func GetUserList3(g *gin.Context) {
	// u := userDao.GetUserList()
	result.UError(g, "没找到")
}

func GetUserList4(g *gin.Context) {
	// u := userDao.GetUserList()
	result.Error(g, err.ErrBind)
}
