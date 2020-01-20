package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	
	uConfig "AppServer/src/config"
	"AppServer/src/dbTransaction"
	"github.com/valyala/fasthttp"
)

var listener net.Listener
var http2Mux *http.ServeMux

var err		  error

var quitHandler chan struct{}

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