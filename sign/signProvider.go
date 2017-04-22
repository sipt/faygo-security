package sign

//ISignProvider sign provider
type ISignProvider interface {
	//GetClientSecret get client secret by clientID.
	GetClientSecret(clientID string) (*ClientInfo, error)
}
