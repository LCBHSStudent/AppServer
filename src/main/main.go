package main

import (
	uConfig "AppServer/src/config"
	"crypto/tls"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

var listener net.Listener
var http2Mux *http.ServeMux

var err		  error

var quitHandler chan struct{}

func main() {
	quitHandler = make(chan struct{}, 1)
	http2Mux = new(http.ServeMux)

	if _, err = tls.LoadX509KeyPair(uConfig.PemPath, uConfig.KeyPath); err != nil {
		os.Exit(uConfig.ErrCertInvalid)
	}

	server := &http.Server {
		ReadHeaderTimeout: 		uConfig.ServiceReadTimeout,
		WriteTimeout:	   		uConfig.ServiceWriteTimeout,
		Handler: 				http2Mux,
	}

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