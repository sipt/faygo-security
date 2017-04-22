package sign

import (
	"github.com/sipt/faygo-security/common"
)

var (
	//ErrorMissingClientID missing clientID in params
	ErrorMissingClientID = common.NewError(common.CodeMissingClientID, "missing clientID")

	//ErrorMissingSign missing sign in params
	ErrorMissingSign = common.NewError(common.CodeMissingSign, "missing sign")

	//ErrorInvalidClientID Invalid clientID by provider
	ErrorInvalidClientID = common.NewError(common.CodeInvalidClientID, "invalid clientID")

	//ErrorInvalidSign sign is error
	ErrorInvalidSign = common.NewError(common.CodeInvalidSign, "invalid sign")

	//ErrorSkipSign sign skip for test
	ErrorSkipSign = common.NewError(common.CodeInvalidSign, "skip sign")
)
