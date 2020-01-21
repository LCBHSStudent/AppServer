package main

import (
	"context"
	"fmt"
	
	"AppServer/src/dbTransaction"
	"AppServer/src/requestHandle"
)

var quitHandler     *context.Context
//var quitHandler chan struct{}

func main() {
	//quitHandler = make(chan struct{}, 1)
	quitHandler  = new(context.Context)
	*quitHandler = context.Background()
	
	fmt.Println(quitHandler)
	
	go dbTransaction.InitDatabaseMySql()
	go dbTransaction.InitDatabaseRedis()
	
	requestHandle.SetMainThreadQuitContext(quitHandler)
	
	go func() {
		requestHandle.InitH2Server()
		requestHandle.CreateH2Listener()
	}()
	
	for {
		select {
		case <- (*quitHandler).Done():
			// do some clear transaction here ...
			break
		default:
			break
		}
	}
	//<- quitHandler
}