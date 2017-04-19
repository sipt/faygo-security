package common

import "fmt"

const (
	//CodeSignFormater fromat error code to string
	CodeSignFormater = "S%05d"

	// <======Sign error code======>

	//CodeNotFoundClientID not found clientID error code
	CodeNotFoundClientID = iota
	//CodeNotFoundSign not found sign error code
	CodeNotFoundSign
	//CodeInvalidClientID invalid clientID error code
	CodeInvalidClientID
	//CodeInvalidSign invalid sign error code
	CodeInvalidSign

	// <======Protect error code======>

	//CodeReplayAttack replay attack error code
	CodeReplayAttack
	//CodeTimestampFormat timstamp format not use RCF3339
	CodeTimestampFormat
)

//ErrorFormater error formater
//type func(string, string)
var ErrorFormater = func(code string, msg string) string {
	return fmt.Sprintf("{\"code\":\"%s\", \"msg\":\"%s\"}", code, msg)
}

//BaseError base error
type BaseError struct {
	Msg  string
	Code int
}

//GetCode get error code and format it
func (b *BaseError) GetCode() string {
	return fmt.Sprint(CodeSignFormater, b.Code)
}

//Error return format error string
func (b *BaseError) Error() string {
	return ErrorFormater(b.GetCode(), b.Msg)
}
