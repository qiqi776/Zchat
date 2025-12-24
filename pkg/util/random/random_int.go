package random

import (
	"math"
	"math/rand/v2"
	"strconv"
	"time"
)
// GetRandomInt 生成指定长度的随机数字
func GetRandomInt(len int) int {
	if len <= 0 {
		return 0
	}
    if len > 18 {
        len = 18
    }
	base := int(math.Pow(10, float64(len-1)))
	return rand.IntN(9*base) + base
}

// GetNowAndLenRandomString 生成 "YYYYMMDD + 随机数" 格式字符串
func GetNowAndLenRandomString(len int) string {
	return time.Now().Format("20060102") + strconv.Itoa(GetRandomInt(len))
}