package login

import (
	"cat-slave/pkg/http/result"
	"cat-slave/pkg/token"
	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func Login(g *gin.Context) {
	// u := userDao.GetUserList()
	token, err := token.Sign(jwt.MapClaims{})
	if err != nil {
		result.Error(g, err)
	} else {
		result.Success(g, map[string]interface{}{
			"token": token,
		})
	}
}
