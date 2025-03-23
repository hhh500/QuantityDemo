package defineModel

import "fmt"

type MyKlineArrayReq struct {
	AcType         string //账户类型
	SymbolName     string //交易对
	KlineInterval  string //k线周期
	KlineSize      int    //k线根数
	StartTimeStamp int64  //k线开始时间
	EndTimeStamp   int64  //k线结束时间
}

//  =============kline通用返回参数=================

type MyKlineArray struct {
	ArrOpenTimeStr     []string  //btcusdt 5m 2021-01-01 00:00:00
	ArrOpenTime        []int64   //开盘毫秒级时间戳
	ArrEndTime         []int64   //结束毫秒级时间戳
	ArrTradeNumber     []int64   //成交笔数
	ArrOpen            []float64 //开盘价
	ArrHigh            []float64 //最高价
	ArrLow             []float64 //最低价
	ArrClose           []float64 //收盘价
	ArrVolume          []float64 //这根K线期间成交量
	ArrQty             []float64 //这根K线期间成交额
	ArrTakerBuyVolume  []float64 //主动买入的成交量
	ArrTakerBuyQty     []float64 //主动买入的成交额
	ArrTakerSellVolume []float64 //主动卖出的成交量
	ArrTakerSellQty    []float64 //主动卖出的成交额
}

func NewMyKlineArrayLen(len int) *MyKlineArray {
	return &MyKlineArray{
		ArrOpenTimeStr:     make([]string, len),
		ArrOpenTime:        make([]int64, len),
		ArrEndTime:         make([]int64, len),
		ArrTradeNumber:     make([]int64, len),
		ArrOpen:            make([]float64, len),
		ArrHigh:            make([]float64, len),
		ArrLow:             make([]float64, len),
		ArrClose:           make([]float64, len),
		ArrVolume:          make([]float64, len),
		ArrQty:             make([]float64, len),
		ArrTakerBuyVolume:  make([]float64, len),
		ArrTakerBuyQty:     make([]float64, len),
		ArrTakerSellVolume: make([]float64, len),
		ArrTakerSellQty:    make([]float64, len),
	}
}

func NewMyKlineArrayCap(len int) *MyKlineArray {
	return &MyKlineArray{
		ArrOpenTimeStr:     make([]string, 0, len),
		ArrOpenTime:        make([]int64, 0, len),
		ArrEndTime:         make([]int64, 0, len),
		ArrTradeNumber:     make([]int64, 0, len),
		ArrOpen:            make([]float64, 0, len),
		ArrHigh:            make([]float64, 0, len),
		ArrLow:             make([]float64, 0, len),
		ArrClose:           make([]float64, 0, len),
		ArrVolume:          make([]float64, 0, len),
		ArrQty:             make([]float64, 0, len),
		ArrTakerBuyVolume:  make([]float64, 0, len),
		ArrTakerBuyQty:     make([]float64, 0, len),
		ArrTakerSellVolume: make([]float64, 0, len),
		ArrTakerSellQty:    make([]float64, 0, len),
	}
}

func (s *MyKlineArray) PrintClose(symbol string) {
	fmt.Printf("time:%d open:%d high:%d low:%d close:%d tradeNumber:%d\n", len(s.ArrClose), len(s.ArrOpen), len(s.ArrHigh), len(s.ArrLow), len(s.ArrClose), len(s.ArrTradeNumber))
	for k, v := range s.ArrClose {
		fmt.Printf("%s %s O:%.4f H:%.4f L:%.4f C:%.4f V:%.2f %d\n",
			symbol, s.ArrOpenTimeStr[k], s.ArrOpen[k], s.ArrHigh[k], s.ArrLow[k], v, s.ArrVolume[k], s.ArrTradeNumber[k])
	}
}
