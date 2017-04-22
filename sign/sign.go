package sign

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
)

const (
	//NameClientID the name of  clientID
	NameClientID = "clientID"
	//NameClientSecret the name of clientSecret
	NameClientSecret = "clientSecret"
	//NameSign the name of sign
	NameSign = "sign"
)

//CheckSign check sign validity
func CheckSign(params map[string]string, provider ISignProvider) (interface{}, bool, error) {
	//check request sign exist
	clientID := params[NameClientID]
	if stringIsEmpty(clientID) {
		return nil, false, ErrorMissingClientID
	}
	clientInfo, err := provider.GetClientSecret(clientID)
	if err != nil {
		return nil, false, err
	}
	if clientInfo == nil {
		return nil, false, ErrorInvalidClientID
	}
	if clientInfo.IsTest {
		return nil, true, nil
	}
	clientSign := params[NameSign]
	if stringIsEmpty(clientSign) {
		return nil, false, ErrorMissingSign
	}
	//sign other params
	clientData, sign, err := Sign(params, clientInfo)
	if err != nil {
		return nil, false, err
	}
	//check request sign invalidity
	if clientSign == sign {
		return clientData, true, nil
	}
	return nil, false, ErrorInvalidSign
}

//Sign sign all params
func Sign(params map[string]string, clientInfo *ClientInfo) (interface{}, string, error) {
	//check params integrity
	clientID := params[NameClientID]
	if stringIsEmpty(clientID) {
		return nil, "", ErrorMissingClientID
	}
	if stringIsEmpty(clientInfo.Secret) {
		return nil, "", ErrorInvalidClientID
	}
	//take out keys
	keys := make([]string, 0)
	var k string
	for k, v := range params {
		if inclueName(k) && v != "" {
			keys = append(keys, k)
		}
	}
	//sort
	sort.Sort(sort.StringSlice(keys))
	signStr := ""
	//splice string  as "k1=v1&k2=v2&..."
	for _, k = range keys {
		if signStr != "" {
			signStr += "&"
		}
		signStr += fmt.Sprint(k, "=", params[k])
	}
	//splice clientSecret to signStr
	signStr += fmt.Sprint("&key=", clientInfo.Secret)
	sign, err := md5Encode(signStr)
	fmt.Println("sign before:", signStr)
	fmt.Println("sign:", sign)
	return clientInfo, sign, err
}

func md5Encode(src string) (string, error) {
	//md5
	hash := md5.New()
	hash.Write([]byte(src))
	cipherText := hash.Sum(nil)
	hexText := make([]byte, 32)
	hex.Encode(hexText, cipherText)
	return string(hexText), nil
}

func stringIsEmpty(v interface{}) bool {
	return v == nil || v == ""
}

func inclueName(param string) bool {
	switch param {
	case NameClientID:
		fallthrough
	case NameSign:
		return false
	default:
		return true
	}
}
