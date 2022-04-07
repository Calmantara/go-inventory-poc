package model

type CommonModel struct {
	ResponseMessage string `json:"message"`
	ResponseType    string `json:"type"`
	ResponseCode    string `json:"code"`
}

type CommonResponse struct {
	CommonModel
	ResponseData interface{} `json:"data,omitempty"`
}

type CommonErrorResponse struct {
	CommonModel
	InvalidArgs interface{} `json:"invalid_args,omitempty"`
}

type CommonResponseType struct {
	HttpCode       int
	CommonResponse CommonResponse
}

type CommonErrorResponseType struct {
	HttpCode            int
	CommonErrorResponse CommonErrorResponse
}
