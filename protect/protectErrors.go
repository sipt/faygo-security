package protect

import "errors"

var (
	//ErrorReplayAttack replay attack
	ErrorReplayAttack = errors.New("request invalid")

	//ErrorTimestampFormat timestamp format not use RCF3339
	ErrorTimestampFormat = errors.New("request params timestamp invalid")
)
