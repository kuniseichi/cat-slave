package user

import (
	userDao "cat-slave/model/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(r *gin.Context) {
	u := userDao.GetUserList()

	r.JSON(http.StatusOK, u)
}
