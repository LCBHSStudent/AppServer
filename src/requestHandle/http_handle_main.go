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

var (
	_listenerList       []*net.Listener
	_listenerCnt        uint = 0

	_http2MuxList       []*http.ServeMux
	_http2MuxCnt        uint = 0
	
	_serverList         []*http.Server
	_serverCnt          uint = 0
	
	//add server/Listener context list
	
	//_mainCtx            context.Context
	_mainCancel         context.CancelFunc
)

func SetMainContextCancelFunc(cancelFunc context.CancelFunc) {
	_mainCancel = cancelFunc
}

func CreateServerAndListener() {
	CreateH2Server()
	CreateH2Listener()
}

func CreateH2Server() {
	var err error
	
	_http2MuxList = append(_http2MuxList, new(http.ServeMux))
	_serverList   = append(_serverList,   new(http.Server))
	
	_http2MuxCnt++; _serverCnt++
	
	if _, err = tls.LoadX509KeyPair(uconfig.PemPath, uconfig.KeyPath); err != nil {
		log.Printf("Load https certificate failed.")
		panic(err)
	}
	
	_serverList[0] = &http.Server {
		ReadHeaderTimeout: 		uconfig.ServiceReadTimeout,
		WriteTimeout:	   		uconfig.ServiceWriteTimeout,
		Handler: 				_http2MuxList[0],
	}
}

func CreateH2Listener() {
	var err error
	
	_listenerList = append(_listenerList, new(net.Listener))
	_listenerCnt++
	
	if *_listenerList[0], err = net.Listen("tcp", ":" +
		strconv.Itoa(uconfig.ServicePort)); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Start listening on https://localhost:" + strconv.Itoa(uconfig.ServicePort))
	}
	
	go _serverList[0].ServeTLS(*_listenerList[0], uconfig.PemPath, uconfig.KeyPath)
}

func CleanUpAllServerAndListener() {
	//for {}  for {}
}

/*
func main() {

	dbTransaction.InitDatabaseMySql()

	quitHandler = make(chan struct{}, 1)
	http2Mux = new(http.ServeMux)

	if err = initEveryThing(); err != nil {
		log.Fatal(err)
	}

	if _, err = tls.LoadX509KeyPair(uConfig.PemPath, uConfig.KeyPath); err != nil {
		os.Exit(uConfig.ErrCertInvalid)
	}

	server := &http.Server {
		ReadHeaderTimeout: 		uConfig.ServiceReadTimeout,
		WriteTimeout:	   		uConfig.ServiceWriteTimeout,
		Handler: 				http2Mux,
	}
	defer server.Close()

	if listener, err = net.Listen("tcp", ":" +
		strconv.Itoa(uConfig.ServicePort)); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Start listening on https://localhost:" + strconv.Itoa(uConfig.ServicePort))
	}

	go server.ServeTLS(listener, uConfig.PemPath, uConfig.KeyPath)

	<-quitHandler
}

func httpHandle(ctx *fasthttp.RequestCtx) {
	body := ctx.Request.Body()

	fmt.Println(string(body))

	ctx.Response.AppendBodyString("ok")
	ctx.Response.SetStatusCode(204)
}

func initEveryThing() error {
	//if err := uConfig.LocalConfigLoader(); err != nil {
	//	return errors.New("local config error")
	//}
	return nil
}

/* http1 to http2 use interface

func getHttpMux() (httpMux, http2Mux *http.ServeMux) {

	httpMux = http.NewServeMux()
	http2Mux = http.NewServeMux()

	x := make(map[string]handlerFunc, 0)
	x["/"] = Home
	x["/v1"] = Hello1

	for k, v := range x {
		redirectURL := http.RedirectHandler(_HTTP2URLBase+k, 307)
		httpMux.Handle(k, redirectURL)
		http2Mux.HandleFunc(k, v)
	}

	return
}

*/