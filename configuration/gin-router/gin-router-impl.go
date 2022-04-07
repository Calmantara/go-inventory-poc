package ginrouter

import (
	"fmt"

	ginroutergroup "github.com/Calmantara/go-inventory-poc/configuration/gin-router-group"
	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	Gin  *gin.Engine
	Port string
}

type Option func(gr *GinRouter)

func NewGinRouter() Router {
	return &GinRouter{
		Gin: gin.New(),
	}
}

func (gr *GinRouter) GROUP(groupPath string, handlers ...gin.HandlerFunc) ginroutergroup.RouterGroup {
	group := gr.Gin.Group(groupPath, handlers...)

	return ginroutergroup.NewGinRouterGroup(group)
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

func (gr *GinRouter) SERVE(ops ...Option) {
	// iterating ops function
	for _, v := range ops {
		v(gr)
	}
	fmt.Printf("Running application in %v", gr.Port)
	gr.Gin.Run(":" + gr.Port)
}

func WithPort(port string) Option {
	return func(gr *GinRouter) {
		gr.Port = port
	}
}
