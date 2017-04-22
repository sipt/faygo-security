package common

import "fmt"

const (
	//CodeSignFormater fromat error code to string
	CodeSignFormater = "S%05d"

	// <======Sign error code======>

	//CodeMissingClientID not found clientID error code
	CodeMissingClientID = iota
	//CodeMissingSign not found sign error code
	CodeMissingSign
	//CodeInvalidClientID invalid clientID error code
	CodeInvalidClientID
	//CodeInvalidSign invalid sign error code
	CodeInvalidSign

	// <======Protect error code======>

	//CodeReplayAttack replay attack error code
	CodeReplayAttack
	//CodeTimestampFormat timstamp format not use RCF3339
	CodeTimestampFormat
	//CodeRequestExpired request expired
	CodeRequestExpired
)

//IError custom error interface
type IError interface {
	GetCode() string
	GetMsg() string
}

//BaseError base error
type BaseError struct {
	msg         string
	code        int
	description string
}

//NewError create a error
func NewError(code int, msg string) error {
	return &BaseError{
		msg:  msg,
		code: code,
	}
}

//GetCode get error code and format it
func (b *BaseError) GetCode() string {
	return fmt.Sprintf(CodeSignFormater, b.code)
}

//GetMsg get error msg
func (b *BaseError) GetMsg() string {
	return b.msg
}

//Error return format error string
func (b *BaseError) Error() string {
	return ErrorFormater(b)
}
