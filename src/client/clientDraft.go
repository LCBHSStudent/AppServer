package client

import (
	uConfig "AppServer/src/config"
	"crypto/tls"
	"fmt"
	"net/http"
)

func InitClientCertificate() {
	transport := &http.Transport{
		TLSClientConfig: 	  &tls.Config{InsecureSkipVerify: true,},
		MaxIdleConns: 		  uConfig.GatewayMaxConnection,
		MaxIdleConnsPerHost:  uConfig.GatewayMaxConnection,
		IdleConnTimeout:      uConfig.GatewayIdleTimeout,
	}
	fmt.Println(transport)
}
