package ginrouter

import (
	ginroutergroup "github.com/Calmantara/go-inventory-poc/configuration/gin-router-group"
	"github.com/gin-gonic/gin"
)

type Router interface {
	GROUP(groupPath string, handlers ...gin.HandlerFunc) ginroutergroup.RouterGroup
	USE(middleware ...gin.HandlerFunc)

	GET(path string, handler ...gin.HandlerFunc)
	POST(path string, handler ...gin.HandlerFunc)
	PUT(path string, handler ...gin.HandlerFunc)
	// DELETE(path string, handler ...gin.HandlerFunc)

	SERVE(ops ...Option)
}
