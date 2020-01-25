package common

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestCrypt(t *testing.T) {
	log.Println("Uint crypt test")
	if err := Crypt(); err != nil {
		t.Fatal(err)
	} else {
		t.Log("\nPass")
	}
}

func TestSubTimeStamp(t *testing.T) {
	t1 := time.Now()
	fmt.Println(t1)
	<-time.After(time.Second)
	
	t2 := time.Now()
	fmt.Println(t2)
	
	fmt.Println(SubTimeStamp(t1, t2))
}

func TestTimeStamp(t *testing.T) {
	// Add 时间相加
	now := time.Now()
	// ParseDuration parses a duration string.
	// A duration string is a possibly signed sequence of decimal numbers,
	// each with optional fraction and a unit suffix,
	// such as "300ms", "-1.5h" or "2h45m".
	//  Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// 10分钟前
	m, _ := time.ParseDuration("-1m")
	m1 := now.Add(m)
	fmt.Println(m1)
	
	// 8个小时前
	h, _ := time.ParseDuration("-1h")
	h1 := now.Add(8 * h)
	fmt.Println(h1)
	
	// 一天前
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println(d1)
	
	PrintSplit(50)
	
	// 10分钟后
	mm, _ := time.ParseDuration("1m")
	mm1 := now.Add(mm)
	fmt.Println(mm1)
	
	// 8小时后
	hh, _ := time.ParseDuration("1h")
	hh1 := now.Add(hh)
	fmt.Println(hh1)
	
	// 一天后
	dd, _ := time.ParseDuration("24h")
	dd1 := now.Add(dd)
	fmt.Println(dd1)
	
	PrintSplit(50)
	
	// Sub 计算两个时间差
	subM := now.Sub(m1)
	fmt.Println(subM.Minutes(), "分钟")
	
	sumH := now.Sub(h1)
	fmt.Println(sumH.Hours(), "小时")
	
	sumD := now.Sub(d1)
	fmt.Printf("%v 天\n", sumD.Hours()/24)
}