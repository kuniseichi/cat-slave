package router

import (
	"cat-slave/handler/passage"
	"cat-slave/handler/remind"
	"cat-slave/handler/sd"
	"cat-slave/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	// By default gin.DefaultWriter = os.Stdout
	// g.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	// your custom format
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	// }))
	//const version = "/api_v1"
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)


	g.NoRoute(func(g *gin.Context) {
		g.String(http.StatusNotFound, "The incorrect API route.")
	})

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	//l := g.Group("login")
	//{
	//	l.GET("", login.Login)
	//}

	u := g.Group("/user")
	// u.Use(middleware.AuthMiddleware())
	{
		u.GET("/wxRecall", remind.Page)
	}

	p := g.Group("/index")
	{
		p.GET("/:keyword", passage.Index)
	}

	r := g.Group("/remind")
	//r.Use(middleware.AuthMiddleware())
	{
		r.GET("")// 获取当前用户当天的事项
		r.POST("") // 新增事项

	}


	return g
}
