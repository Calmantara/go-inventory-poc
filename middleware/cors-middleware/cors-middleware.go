package corsmiddleware

import "github.com/gin-gonic/gin"

type CorsMiddlware interface {
	CorsRequest(ctx *gin.Context)
}
