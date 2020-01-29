package requestHandle

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"strconv"
	
	uconfig "AppServer/src/config"
)

type handlerFunc func(
	w http.ResponseWriter, r *http.Request,
)

var (
	_listenerList       []*net.Listener
	_listenerCnt        uint = 0

	_http2MuxList       []*http.ServeMux
	_http2MuxCnt        uint = 0
	
	_serverList         []*http.Server
	_serverAddr         []string
	_serverCnt          uint = 0
	
	
	//add server/Listener context list
	
	//_mainCtx            context.Context
	_mainCancel         context.CancelFunc
)

func SetMainContextCancelFunc(cancelFunc context.CancelFunc) {
	_mainCancel = cancelFunc
}

func CreateServerAndListener() {
	createH2Server()
	createH2Listener()
}

func createH2Server() {
	var err error
	
	_http2MuxList = append(_http2MuxList, new(http.ServeMux))
	_serverList   = append(_serverList,   new(http.Server))
	
	
	
	if _, err = tls.LoadX509KeyPair(uconfig.PemPath, uconfig.KeyPath); err != nil {
		log.Printf("Load https certificate failed.")
		panic(err)
	}
	
	setHandleStrategy(_http2MuxList[0])
	
	_serverList[_serverCnt] = &http.Server {
		ReadHeaderTimeout: 		uconfig.ServiceReadTimeout,
		WriteTimeout:	   		uconfig.ServiceWriteTimeout,
		Handler: 				_http2MuxList[_http2MuxCnt],
	}
	
	_http2MuxCnt++; _serverCnt++
}

func createH2Listener() {
	var err error
	
	_listenerList = append(_listenerList, new(net.Listener))
	
	
	if *_listenerList[_listenerCnt], err = net.Listen("tcp", ":" +
		strconv.Itoa(uconfig.ServicePort)); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Start listening on https://localhost:" + strconv.Itoa(uconfig.ServicePort))
	}
	
	_listenerCnt++
	
	go _serverList[0].ServeTLS(*_listenerList[0], uconfig.PemPath, uconfig.KeyPath)
}

func setHandleStrategy(mux *http.ServeMux) {
	serveContent := make(map[string]handlerFunc, 0)
	
	serveContent["/"] = testMain
	
	for path, fun := range serveContent {
		//redirectURL := http.RedirectHandler("https://localhost:"+strconv.Itoa(uconfig.ServicePort)+
		//	path, 307)
		//mux.Handle(path, redirectURL)
		mux.HandleFunc(path, fun)
	}
}

func CleanUpAllServerAndListener() {
	//for {}  for {}
}