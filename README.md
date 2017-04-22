# gin-security
Gin security middleware

## Sign
### 介绍
用于验证接口请求完整性，对请求的所有KV（除媒体文件）以及时间戳`timestamp`和随机字符串`nonce`，进行大小写排序，用`&`、`=`相连，最后连接上`key=secret`,进行`MD5`。服务端对此进行验证。
### 使用
Sign需要实现`ISignProvider`, `GetClientSecret`通过`clientID`来获取`ClientInfo`
``` go
//ISignProvider sign provider
type ISignProvider interface {
	//GetClientSecret get client secret by clientID.
	GetClientSecret(clientID string) (*ClientInfo, error)
}
```
`CheckSign`用于获验证签名是否正确，`Sign`用于生成签名
example
``` go
//GetClientSecret get client secret by clientID.
func (SignProviderImpl) GetClientSecret(clientID string) (*sign.ClientInfo, error) {
	conn := r.GetConn()
	bytes, err := redis.Bytes(conn.Do(r.GET, joinKey(clientID)))
	if err != nil {
		return nil, err
	}
	var clientInfo = new(sign.ClientInfo)
	err = proto.Unmarshal(bytes, clientInfo)
	if err != nil {
		return nil, errors.New("clientID invalid")
	}
	return clientInfo, nil
}
```
``` go
captcha.GET("/", sign.NewSignMiddleware(SignProvider), handler.CaptchaFactory)
```

## Protect
### 介绍
对请求中的`timestamp`和`nonce`进行验证，`timestamp`和`nonce`唯一标识一个请求，当出现请求重复时，认为是请求重放攻击，可以自行处理，如：下次发起请求时必须带有验证码。当`timestamp`超出了预设时间，如60s，就返回错误信息并且带上服务器当前时间(RFC 3339)，再次请求，超过指定次数就要求携带验证码。
### 使用
实现下面的接口，`IsExist`检查请求唯一性， `DealWithError`处理错误，如：`ErrorReplayAttac`请求重放,`ErrorRequestExpired`请求过期 等等。
``` go
//IProtectProvider protect provider
type IProtectProvider interface {
	//IsExist judge request isExist
	IsExist(timestamp, nonce string) (bool, error)

	//DealWithError deal with error
	//error : ErrorReplayAttack, ErrorTimestampFormat, ErrorRequestExpired
	DealWithError(ctx *gin.Context, err error)
}
```
``` go
captcha.GET("/", protect.NewProtectMiddleware(ProtectProvider), handler.CaptchaFactory)
```

