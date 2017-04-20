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
			provider.DealWithError(ctx, err)
			ctx.JSON(422, common.Pack(ctx,
				gin.H{
					"timestamp": time.Now().Format(time.RFC3339),
				}, err))
		}
		ctx.Next()
	}
}
