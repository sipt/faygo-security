package sign

import "errors"

var (
	//ErrorMissingClientID missing clientID in params
	ErrorMissingClientID = errors.New("missing clientID")

	//ErrorMissingSign missing sign in params
	ErrorMissingSign = errors.New("missing sign")

	//ErrorInvalidClientID Invalid clientID by provider
	ErrorInvalidClientID = errors.New("invalid clientID")

	//ErrorInvalidSign sign is error
	ErrorInvalidSign = errors.New("invalid sign")
)
