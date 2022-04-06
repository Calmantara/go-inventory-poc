package ginroutergroup

import "github.com/gin-gonic/gin"

type GinRouterGroup struct {
	Rg *gin.RouterGroup
}

func NewGinRouterGroup(rg *gin.RouterGroup) RouterGroup {
	return &GinRouterGroup{
		Rg: rg,
	}
}

func (gr *GinRouterGroup) GET(path string, handler ...gin.HandlerFunc) {
	gr.Rg.GET(path, handler...)
}

func (gr *GinRouterGroup) POST(path string, handler ...gin.HandlerFunc) {
	gr.Rg.POST(path, handler...)
}

func (gr *GinRouterGroup) PUT(path string, handler ...gin.HandlerFunc) {
	gr.Rg.PUT(path, handler...)
}

func (gr *GinRouterGroup) DELETE(path string, handler ...gin.HandlerFunc) {
	gr.Rg.DELETE(path, handler...)
}
