package protect

//IProtectProvider protect provider
type IProtectProvider interface {
	//IsExist judge request isExist
	IsExist(timestamp, nonce string) (bool, error)
}
