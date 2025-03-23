package apiImplBn

import (
	"QuantityDemo/api/exchange/exchangeImplBn/bnSdk"
	"QuantityDemo/api/exchange/exchangeImplBn/constantBn"
)

var MyRestImplBn = NewRestImplBn("", "", "")

type RestApiImplBn struct {
	accountKey string
	ApiKey     string
	secretKey  string
	future     *bnSdk.FutureRestClient
}

func NewRestImplBn(accountKey, apiKey, secretKey string) *RestApiImplBn {
	return &RestApiImplBn{
		accountKey: accountKey,
		ApiKey:     apiKey,
		secretKey:  secretKey,
		future:     bnSdk.NewFutureRestClient(constantBn.FUTURE, apiKey, secretKey),
	}
}
