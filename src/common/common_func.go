package common

import (
	"log"
)

func IF(condition bool, trueVal interface{}, falseVal interface{}) {
	if condition {
		log.Println(trueVal)
	} else {
		log.Println(falseVal)
	}
}
