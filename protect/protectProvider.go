package protect

import "github.com/gin-gonic/gin"

//IProtectProvider protect provider
type IProtectProvider interface {
	//IsExist judge request isExist
	IsExist(timestamp, nonce string) (bool, error)

	//DealWithError deal with error
	//error : ErrorReplayAttack, ErrorTimestampFormat, ErrorRequestExpired
	DealWithError(ctx *gin.Context, err error)
}
