package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//ErrorFormater error formater
//type func(string, string)
var ErrorFormater = func(b *BaseError) string {
	return fmt.Sprintf("{\"code\":\"%s\", \"msg\":\"%s\"}", b.GetCode(), b.msg)
}

//BaseResult result formater
type BaseResult struct {
	APIVersion string
	ID         string
	Method     string
	Error      interface{}
	Data       interface{}
}

//ErrorResult error formater in result json
type ErrorResult struct {
	Code string
	Msg  string
}

//M alias of map[string]interface{}
type M map[string]interface{}

//Pack pack data or error to map for result to client
//eg. Pack(ctx, data, err)
//    Pack(ctx, err)
//    Pack(ctx, data)
var Pack = func(ctx *gin.Context, params ...interface{}) interface{} {
	flag := 0
	result := &BaseResult{
		Method: ctx.Request.Method,
	}
	for _, p := range params {
		if p == nil {
			continue
		}
		if e, ok := p.(IError); ok && flag&1 == 0 {
			result.Error = &ErrorResult{
				Code: e.GetCode(),
				Msg:  e.GetMsg(),
			}
		} else if e, ok := p.(error); ok && flag&1 == 0 {
			result.Error = &ErrorResult{
				Msg: e.Error(),
			}
		} else if flag&2 == 1 {
			result.Data = p
		}
	}
	return result
}
