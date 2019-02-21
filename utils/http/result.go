package result

import (
	"net/http"

	"cat-slave/utils/err"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// func SendResponse(c *gin.Context, err error, data interface{}) {
// 	code, message := errno.DecodeErr(err)

// 	// always return http.StatusOK
// 	c.JSON(http.StatusOK, Response{
// 		Code:    code,
// 		Message: message,
// 		Data:    data,
// 	})
// }

func Success(g *gin.Context, data map[string]interface{}) {
	g.JSON(http.StatusOK, Result{
		Success: true,
		Code:    err.OK.Code,
		Message: err.OK.Message,
		Data:    data,
	})
}

func UError(g *gin.Context, msg string) {
	var message string
	if msg == "" {
		message = msg
	} else {
		message = err.Fail.Message
	}
	g.JSON(http.StatusOK, Result{
		Success: false,
		Code:    err.Fail.Code,
		Message: message,
		Data:    nil,
	})
}
func Error(g *gin.Context, errno *err.Errno) {

	g.JSON(http.StatusOK, Result{
		Success: false,
		Code:    errno.Code,
		Message: errno.Message,
		Data:    nil,
	})
}
