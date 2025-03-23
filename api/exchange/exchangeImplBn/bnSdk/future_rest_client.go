package bnSdk

type FutureRestClient struct {
	rest *RestClient
}

func NewFutureRestClient(acType, apiKey, secretKey string) *FutureRestClient {
	return &FutureRestClient{
		rest: NewRestClient(acType, apiKey, secretKey),
	}
}
