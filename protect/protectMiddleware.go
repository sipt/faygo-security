package protect

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sipt/faygo-security/common"
)

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
			common.Error(err.Error())
			provider.DealWithError(ctx, err)
			var data interface{}
			if err == ErrorRequestExpired {
				data = gin.H{
					"timestamp": time.Now().Format(time.RFC3339),
				}
			}
			ctx.JSON(422, common.Pack(ctx, data, err))
		}
		ctx.Next()
	}
}
