package commonhandler

import (
	"github.com/gin-gonic/gin"
)

type CommonHandler interface {
	CommonErrorResponseBuilder(ctx *gin.Context, eType ErrorType, errorArgs interface{})
	CommonResponseBuilder(ctx *gin.Context, rType ResponseType, data interface{})
}
