package dbTransaction

import (
	"log"
	"strconv"
	"testing"
	"time"
)

func TestInitDatabaseRedis(t *testing.T) {
	InitDatabaseRedis()
}

func TestTimedQuery(t *testing.T) {
	//timer1 := time.NewTimer(time.Second * 5)
	
	for i := 0; i < 5; i++ {
		select {
		case <-time.After(time.Second * 2):
			log.Println("start No.", i+1, "query")
			testUser := strconv.Itoa(i+1)
			
			setUserToken(testUser)
			log.Println(testUser)
			// test query
		}
	}
}

func TestCheckUserToken(t *testing.T) {
	CheckUserToken("1")
}