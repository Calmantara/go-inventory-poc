package commonhandler

import (
	"net/http"

	"github.com/Calmantara/go-inventory-poc/model"
	"github.com/gin-gonic/gin"
)

// custom data type and const
type ResponseType string
type ErrorType string

const (
	INTERNAL     = ErrorType("INTERNAL_SERVER_ERROR")
	BAD_REQUEST  = ErrorType("BAD_REQUEST")
	UNAUTHORIZED = ErrorType("UNAUTHORIZED")
	NOT_FOUND    = ErrorType("NOT_FOUND")

	SUCCESS = ResponseType("SUCCESS")
)

type CommonHandlerImpl struct {
}

func NewCommonHandler() CommonHandler {
	return &CommonHandlerImpl{}
}

func (c *CommonHandlerImpl) CommonErrorResponseBuilder(ctx *gin.Context, eType ErrorType, errorArgs interface{}) {
	var commonResponse model.CommonErrorResponseType
	// switcher for common error response
	switch eType {
	case BAD_REQUEST:
		commonResponse = model.CommonErrorResponseType{
			HttpCode: http.StatusBadRequest,
			CommonErrorResponse: model.CommonErrorResponse{
				CommonModel: model.CommonModel{
					ResponseMessage: "bad request payload",
					ResponseType:    string(BAD_REQUEST),
					ResponseCode:    "98",
				},
				InvalidArgs: errorArgs,
			},
		}
	case UNAUTHORIZED:
		commonResponse = model.CommonErrorResponseType{
			HttpCode: http.StatusUnauthorized,
			CommonErrorResponse: model.CommonErrorResponse{
				CommonModel: model.CommonModel{
					ResponseMessage: "unauthorized request",
					ResponseType:    string(UNAUTHORIZED),
					ResponseCode:    "97",
				},
				InvalidArgs: errorArgs,
			},
		}
	case NOT_FOUND:
		commonResponse = model.CommonErrorResponseType{
			HttpCode: http.StatusNotFound,
			CommonErrorResponse: model.CommonErrorResponse{
				CommonModel: model.CommonModel{
					ResponseMessage: "entity not found",
					ResponseType:    string(NOT_FOUND),
					ResponseCode:    "96",
				},
				InvalidArgs: errorArgs,
			},
		}
	default:
		commonResponse = model.CommonErrorResponseType{
			HttpCode: http.StatusInternalServerError,
			CommonErrorResponse: model.CommonErrorResponse{
				CommonModel: model.CommonModel{
					ResponseMessage: "internal server error",
					ResponseType:    string(INTERNAL),
					ResponseCode:    "99",
				},
				InvalidArgs: errorArgs,
			},
		}
	}
	// abort context
	ctx.AbortWithStatusJSON(commonResponse.HttpCode, commonResponse.CommonErrorResponse)
}

func (c *CommonHandlerImpl) CommonResponseBuilder(ctx *gin.Context, rType ResponseType, data interface{}) {
	var commonResponse model.CommonResponseType
	// common response switcher
	switch rType {
	default:
		commonResponse = model.CommonResponseType{
			HttpCode: http.StatusOK,
			CommonResponse: model.CommonResponse{
				CommonModel: model.CommonModel{
					ResponseMessage: "request succeed",
					ResponseType:    string(SUCCESS),
					ResponseCode:    "00",
				},
				ResponseData: data,
			},
		}
	}
	// success response
	ctx.JSONP(commonResponse.HttpCode, commonResponse.CommonResponse)
}
