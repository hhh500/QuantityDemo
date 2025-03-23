package exchangeApi

import (
	"QuantityDemo/api/defineModel"
)

// RestApiInterface 交易所api层接口
type RestApiInterface interface {
	GetMyKlineArray(req *defineModel.MyKlineArrayReq) (*defineModel.MyKlineArray, error) //获取K线
}
