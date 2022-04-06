package ginroutergroup

import "github.com/gin-gonic/gin"

type RouterGroup interface {
	GET(path string, handler ...gin.HandlerFunc)
	POST(path string, handler ...gin.HandlerFunc)
	PUT(path string, handler ...gin.HandlerFunc)
	DELETE(path string, handler ...gin.HandlerFunc)
}
