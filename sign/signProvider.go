package sign

//ClientInfo client info for sign provider
type ClientInfo struct {
	ClientSecret string
	Data         interface{}
}

//ISignProvider sign provider
type ISignProvider interface {
	//GetClientSecret get client secret by clientID.
	GetClientSecret(clientID string) (*ClientInfo, error)
}
