package main

import (
	"context"
	"fmt"
	"log"
	
	"AppServer/src/dbTransaction"
	"AppServer/src/requestHandle"
)

var quitHandler     context.Context
var cancelFunc      context.CancelFunc
//var quitHandler chan struct{}

func main() {
	//quitHandler = make(chan struct{}, 1)
	quitHandler, cancelFunc = context.WithCancel(context.Background())
	
	fmt.Println(quitHandler)
	
	go dbTransaction.InitDatabaseMySql()
	go dbTransaction.InitDatabaseRedis()
	
	go func() {
		requestHandle.SetMainContextCancelFunc(cancelFunc)
		requestHandle.CreateServerAndListener()
	}()
	
	select {
	case <- quitHandler.Done():
		log.Println("Main thread will be finished soon.")
		// do some clear transaction here ...
		return
	}
	//<- quitHandler
}