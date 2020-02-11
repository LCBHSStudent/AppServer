package httpRouter

import (
	"net/http"
)
/**
提供基本的路由功能，添加路由，查找路由
*/

const (
	GET         = iota
	POST
	PUT
	DELETE
	CONNECTIBNG
	HEAD
	OPTIONS
	PATCH
	TRACE
)

func NewRouter() MethodMaps {
	return []handler{
		GET:    make(handler),
		POST:   make(handler),
		PUT:    make(handler),
		DELETE: make(handler),
	}
}

type MethodMaps [] handler
type handler map[string]HandlerMapped

//实现IOdServer的接口，以及http提供ServeHttp方法
type OdServer struct {
	router MethodMaps
}

type IOdServer interface {
	GET(url string, f HandlerFunc)
	POST(url string, f HandlerFunc)
	PUT(url string, f HandlerFunc)
	DELETE(url string, f HandlerFunc)
}

type HandlerMapped struct {
	f HandlerFunc
}
//接口函数单位，即我们编写代码逻辑的函数
type HandlerFunc func(w http.ResponseWriter, req *http.Request)