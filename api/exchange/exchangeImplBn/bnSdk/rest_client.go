package bnSdk

import (
	"QuantityDemo/api/exchange/exchangeImplBn/constantBn"
	"QuantityDemo/utils/timeUtils"
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type doFunc func(req *http.Request) (*http.Response, error)

type RestClient struct {
	acType     string
	apiKey     string
	secretKey  string
	baseURL    string
	httpClient *http.Client
	do         doFunc
}

func getRestHost(acType string) string {
	switch acType {
	case constantBn.SPOT:
		return "https://api.binance.com"
	case constantBn.FUTURE:
		return "https://fapi.binance.com"
	case constantBn.SWAP:
		return "https://dapi.binance.com"
	}
	return ""
}

func NewRestClient(acType, apiKey, secretKey string) *RestClient {
	return &RestClient{
		acType:     acType,
		baseURL:    getRestHost(acType),
		httpClient: http.DefaultClient,
	}
}

func (c *RestClient) parseRequest(r *request, opts ...RequestOption) (err error) {
	for _, opt := range opts {
		opt(r)
	}
	// 参数校验
	r.validate()
	// 构建 URL
	fullURL := fmt.Sprintf("%s%s", c.baseURL, r.endpoint)
	// 设置query参数
	if r.receiveWindow > 0 {
		r.setQueryParam(receiveWindowKey, r.receiveWindow)
	}
	if r.secType == secTypeSigned {
		r.setQueryParam(timestampKey, timeUtils.GetNowTimeUnixMilli())
	}
	// 编码 query & form
	queryString := r.query.Encode()
	bodyString := r.form.Encode()
	// 构建请求体
	var body *bytes.Buffer
	if bodyString != "" {
		body = bytes.NewBufferString(bodyString)
	} else {
		body = &bytes.Buffer{}
	}
	// 构建 Header(优先用已有 header)
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}
	header.Set("User-Agent", "binance-connector-go/0.8.0")
	header.Set("Content-Type", "application/json")
	if r.secType == secTypeAPIKey || r.secType == secTypeSigned {
		header.Set("X-MBX-APIKEY", c.apiKey)
	}
	if r.secType == secTypeSigned {
		raw := fmt.Sprintf("%s%s", queryString, bodyString)
		mac := hmac.New(sha256.New, []byte(c.secretKey))
		_, err = mac.Write([]byte(raw))
		if err != nil {
			return err
		}
		v := url.Values{}
		v.Set(signatureKey, fmt.Sprintf("%x", (mac.Sum(nil))))
		if queryString == "" {
			queryString = v.Encode()
		} else {
			queryString = fmt.Sprintf("%s&%s", queryString, v.Encode())
		}
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *RestClient) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	f := c.do
	if f == nil {
		f = c.httpClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {

		cerr := res.Body.Close()
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	return data, nil
}
