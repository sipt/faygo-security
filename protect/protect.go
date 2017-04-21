package protect

import (
	"time"
)

var (
	//GlobalConfig protect config
	GlobalConfig = &Config{
		ValidTime: 60 * time.Second,
	}
)

//Config protect config
type Config struct {
	//ValidTime: request lag (seconds)
	ValidTime time.Duration
}

//CheckRequestValidity check request validity
//@param timestamp: date formatted RFC 3339
func CheckRequestValidity(timestamp, nonce string, provider IProtectProvider) (bool, error) {
	//check request replay attack
	requestTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return false, ErrorTimestampFormat
	}
	now := time.Now()
	if now.Sub(requestTime) > GlobalConfig.ValidTime {
		return false, ErrorRequestExpired
	}
	isExist, err := provider.IsExist(timestamp, nonce)
	if err != nil {
		return false, err
	}
	if isExist {
		return false, ErrorReplayAttack
	}
	return true, nil
}
