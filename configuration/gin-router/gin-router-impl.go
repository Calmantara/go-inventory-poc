package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	Gin *gin.Engine
}

func NewGinRouter() Router {
	return &GinRouter{
		Gin: gin.New(),
	}
}

func (gr *GinRouter) GROUP(groupPath string, handlers ...gin.HandlerFunc) routergroup.RouterGroup {
	group := gr.Gin.Group(groupPath, handlers...)

	return routergroup.NewGinRouterGroup(group)
}

func (gr *GinRouter) USE(middleware ...gin.HandlerFunc) {
	gr.Gin.Use(middleware...)
}

func (gr *GinRouter) GET(path string, handler ...gin.HandlerFunc) {
	gr.Gin.GET(path, handler...)
}

func (gr *GinRouter) POST(path string, handler ...gin.HandlerFunc) {
	gr.Gin.POST(path, handler...)
}

func (gr *GinRouter) PUT(path string, handler ...gin.HandlerFunc) {
	gr.Gin.PUT(path, handler...)
}

func (gr *GinRouter) DELETE(path string, handler ...gin.HandlerFunc) {
	gr.Gin.DELETE(path, handler...)
}

func (gr *GinRouter) SERVE(port string) {
	if port == "" {
		port = "5000"
	}
	fmt.Printf("Running application in %v", port)
	gr.Gin.Run(":" + port)
}
