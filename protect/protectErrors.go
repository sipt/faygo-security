package protect

import "github.com/sipt/faygo-security/common"

var (
	//ErrorReplayAttack replay attack
	//replay request to server
	ErrorReplayAttack = common.NewError(common.CodeReplayAttack, "request invalid")

	//ErrorTimestampFormat timestamp format not use RCF3339
	ErrorTimestampFormat = common.NewError(common.CodeTimestampFormat, "request params timestamp invalid")

	//ErrorRequestExpired request expired
	//request expired
	//The clinet and server have time difference. Server return
	ErrorRequestExpired = common.NewError(common.CodeRequestExpired, "request expired")
)
