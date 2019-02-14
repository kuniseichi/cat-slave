package main

import (
	"cat-slave/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a router without any middleware by default
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load()

}
