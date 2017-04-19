package sign

import (
	"github.com/gin-gonic/gin"
)

const (
	//NameClientData the name of clientData in ctx
	NameClientData = "clientData"
)

//NewSignMiddleware new sign middleware
func NewSignMiddleware(provider ISignProvider) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		params := make(map[string]string)

		//获取取有的参数转成map
		p := ctx.Request.URL.Query()
		for k, vs := range p {
			if vs[0] != "" {
				params[k] = vs[0]
			}
		}
		ctx.Request.ParseForm()
		ctx.Request.ParseMultipartForm(32 << 20)
		p = ctx.Request.PostForm
		for k, vs := range p {
			if vs[0] != "" {
				params[k] = vs[0]
			}
		}
		clientData, _, err := CheckSign(params, provider)
		if err != nil {
			ctx.String(400, err.Error())
		} else {
			ctx.Set(NameClientData, clientData)
			ctx.Next()
		}
	}
}
