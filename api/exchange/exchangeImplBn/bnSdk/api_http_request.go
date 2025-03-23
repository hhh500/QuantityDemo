package bnSdk

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type params map[string]interface{}

type request struct {
	method        string      //http请求方法
	endpoint      string      //请求路径
	query         url.Values  //请求参数
	form          url.Values  //表单参数
	receiveWindow int64       //接收窗口
	secType       secType     //签名类型
	header        http.Header //请求头
	body          io.Reader   //请求体
	fullURL       string      //完整URL
}

func (r *request) addQueryParam(key string, value interface{}) *request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add(key, fmt.Sprintf("%v", value))
	return r
}

func (r *request) setQueryParam(key string, value interface{}) *request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

func (r *request) setFormParam(key string, value interface{}) *request {
	if r.form == nil {
		r.form = url.Values{}
	}
	r.form.Set(key, fmt.Sprintf("%v", value))
	return r
}

func (r *request) setQueryParams(m params) *request {
	for k, v := range m {
		r.setQueryParam(k, v)
	}
	return r
}

// setFormParams set params with key/values to request form body
func (r *request) setFormParams(m params) *request {
	for k, v := range m {
		r.setFormParam(k, v)
	}
	return r
}

func (r *request) validate() {
	if r.query == nil {
		r.query = url.Values{}
	}
	if r.form == nil {
		r.form = url.Values{}
	}
}

type RequestOption func(*request)
