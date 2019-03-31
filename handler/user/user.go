package user

import (
	"cat-slave/model/user"
	"cat-slave/pkg/http/result"
	"github.com/gin-gonic/gin"
)

func GetUserList(g *gin.Context) {
	u := userDao.GetUserList()
	result.Success(g, map[string]interface{}{
		"userList": u,
		"token":    "",
	})

}

