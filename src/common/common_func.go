package common

import "fmt"

func IF(condition bool, trueVal interface{}, falseVal interface{}) {
	if condition {
		fmt.Println(trueVal)
	} else {
		fmt.Println(falseVal)
	}
}
