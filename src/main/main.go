package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	log.Fatal(fasthttp.ListenAndServe(":8081", func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Path())
		switch path {
		case "/post":
			httpHandle(ctx)
		default:
			fmt.Print("————————————")
		}
	}))
}

func httpHandle(ctx *fasthttp.RequestCtx) {
	body := ctx.Request.Body()

	fmt.Println(string(body))

	ctx.Response.AppendBodyString("ok")
	ctx.Response.SetStatusCode(204)
}