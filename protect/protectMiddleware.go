package protect

import "github.com/gin-gonic/gin"

const (
	//NameTimestamp the name of timestamp in ctx.params
	NameTimestamp = "timestamp"
	//NameNonce the name of nonce in ctx.params
	NameNonce = "nonce"
)

//NewProtectMiddleware new protect middleware
func NewProtectMiddleware(provider IProtectProvider) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		timestamp := ctx.Param(NameTimestamp)
		nonce := ctx.Param(NameNonce)
		_, err := CheckRequestValidity(timestamp, nonce, provider)
		if err != nil {
			ctx.String(422, err.Error())
		}
		ctx.Next()
	}
}
