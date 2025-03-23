package timeUtils

import "time"

// 获取当前时间的毫秒级时间戳
func GetNowTimeUnixMilli() int64 {
	return time.Now().UnixMilli()
}
