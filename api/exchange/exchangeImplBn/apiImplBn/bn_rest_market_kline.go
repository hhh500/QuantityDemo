package apiImplBn

import (
	"QuantityDemo/api/defineModel"
	"QuantityDemo/api/exchange/exchangeImplBn/constantBn"
	"errors"
)

func (b *RestApiImplBn) getFutureKlineRest(req *defineModel.MyKlineArrayReq) (kline *defineModel.MyKlineArray, err error) {
	client := b.future.NewFutureKlines().Symbol(req.SymbolName).Interval(req.KlineInterval).Limit(req.KlineSize)
	if req.StartTimeStamp > 0 {
		client.StartTime(req.StartTimeStamp)
	}
	if req.EndTimeStamp > 0 {
		client.EndTime(req.EndTimeStamp)
	}
	return client.Do()
}

func (b *RestApiImplBn) GetMyKlineArray(req *defineModel.MyKlineArrayReq) (*defineModel.MyKlineArray, error) {
	switch req.AcType {
	case constantBn.FUTURE:
		return b.getFutureKlineRest(req)
	default:
		return nil, errors.New("getKlineRest error:acType:" + req.AcType)
	}
}
