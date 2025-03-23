package bnSdk

import (
	"QuantityDemo/api/defineModel"
	"net/http"

	"github.com/tidwall/gjson"
)

type FutureKlinesReq struct {
	Symbol    string `json:"symbol"`    //YES
	Interval  string `json:"interval"`  //YES	详见枚举定义：K线间隔
	StartTime *int64 `json:"startTime"` //NO
	EndTime   *int64 `json:"endTime"`   //NO
	Limit     *int   `json:"limit"`     //NO	默认 500; 最大 1000.
}

type FutureKlinesApi struct {
	client *FutureRestClient
	req    *FutureKlinesReq
}

func (api *FutureKlinesApi) Symbol(symbol string) *FutureKlinesApi {
	api.req.Symbol = symbol
	return api
}

func (api *FutureKlinesApi) Interval(interval string) *FutureKlinesApi {
	api.req.Interval = interval
	return api
}

func (api *FutureKlinesApi) StartTime(startTime int64) *FutureKlinesApi {
	api.req.StartTime = &startTime
	return api
}

func (api *FutureKlinesApi) EndTime(endTime int64) *FutureKlinesApi {
	api.req.EndTime = &endTime
	return api
}

func (api *FutureKlinesApi) Limit(limit int) *FutureKlinesApi {
	api.req.Limit = &limit
	return api
}

// 行情接口
// binance FUTURE FutureKlines restK线数据 (MARKET_DATA)
func (client *FutureRestClient) NewFutureKlines() *FutureKlinesApi {
	return &FutureKlinesApi{
		client: client,
		req:    &FutureKlinesReq{},
	}
}

func (api *FutureKlinesApi) Do(opts ...RequestOption) (*defineModel.MyKlineArray, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/klines",
		secType:  secTypeNone, //不需要签名
	}
	r.setQueryParam("symbol", api.req.Symbol)
	r.setQueryParam("interval", api.req.Interval)
	if api.req.Limit != nil {
		r.setQueryParam("limit", *api.req.Limit)
	}
	if api.req.StartTime != nil {
		r.setQueryParam("startTime", *api.req.StartTime)
	}
	if api.req.EndTime != nil {
		r.setQueryParam("endTime", *api.req.EndTime)
	}
	data, err := api.client.rest.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := gjson.Parse(string(data))
	num := len(result.Array())
	kline := defineModel.NewMyKlineArrayLen(num)
	for index, item := range result.Array() {
		kline.ArrOpenTime[index] = item.Get("0").Int()
		kline.ArrOpen[index] = item.Get("1").Float()
		kline.ArrHigh[index] = item.Get("2").Float()
		kline.ArrLow[index] = item.Get("3").Float()
		kline.ArrClose[index] = item.Get("4").Float()
		volmue := item.Get("5").Float()
		kline.ArrVolume[index] = volmue
		kline.ArrEndTime[index] = item.Get("6").Int()
		qty := item.Get("7").Float()
		kline.ArrQty[index] = qty
		kline.ArrTradeNumber[index] = item.Get("8").Int()
		takerBuyVolume := item.Get("9").Float()
		kline.ArrTakerBuyVolume[index] = takerBuyVolume
		takerBuyQty := item.Get("10").Float()
		kline.ArrTakerBuyQty[index] = takerBuyQty
		kline.ArrTakerSellVolume[index] = volmue - takerBuyVolume
		kline.ArrTakerSellQty[index] = qty - takerBuyQty
	}
	return kline, nil
}
