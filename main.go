package main

import (
	"cat-slave/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化gin
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load()

}
