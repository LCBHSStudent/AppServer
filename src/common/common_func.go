package common

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// 用于模拟三目运算符
// IF(a > b, a, b) 打印日志或调试时使用
func IF(condition bool, trueVal interface{}, falseVal interface{}) {
	if condition {
		log.Println(trueVal)
	} else {
		log.Println(falseVal)
	}
}

// 求两个时间戳的差值，返回int数据, 单位：秒
func SubTimeStamp(formerTime time.Time, laterTime time.Time) float64 {
	sub := laterTime.Sub(formerTime)
	//second := int(sub.Seconds())
	return sub.Seconds()
}

func PrintSplit(count int) {
	print("#")
	fmt.Print(strings.Repeat("-", count))
	print("#\n")
}