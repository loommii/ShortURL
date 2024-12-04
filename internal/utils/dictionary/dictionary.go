package dictionary

import (
	"math/rand"
	"sync/atomic"
)

var Dictionary *dictionary

type dictionary struct {
	Counter uint64   // 计数器
	Data    [62]byte // 乱序字典
	Lan     int      // 输出的字符串长度 目前支持6
}

func init() {
	Dictionary = NewDictionary()
}
func NewDictionary() *dictionary {
	d := dictionary{}
	// d.Counter = rand.Uint64() // 默认随机数字
	d.Lan = 6
	d.Data = [62]byte{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}

	rand.Shuffle(len(d.Data), func(i, j int) { // 乱序
		d.Data[i], d.Data[j] = d.Data[j], d.Data[i]
	})
	return &d
}
func (d *dictionary) Get() string {
	num := atomic.AddUint64(&d.Counter, 1)
	numL := make([]byte, 10)
	for i := len(numL) - 1; i >= 0 && num > 0; i-- {
		numL[i] = byte(num % 10)
		num = num / 10
	}
	for i := len(numL) - 1; i >= 0; i-- {
		num *= 10
		num += uint64(numL[i])
	}
	ans := make([]byte, d.Lan)
	// 通过数字转62进制对应乱序字典
	for i := 0; i < d.Lan; i++ {
		ans[i] = d.Data[num%62]
		num = num / 62
	}
	return string(ans)
}
