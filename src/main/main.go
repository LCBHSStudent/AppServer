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
		// 这就偷懒不做并发了罢
		dbTransaction.DisconnectDatabaseRedis()
		dbTransaction.DisconnectMySql()
		// 清理现场，释放句柄和内存
		// ...
		return
	}
	//<- quitHandler
}