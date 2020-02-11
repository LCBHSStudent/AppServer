package requestHandle

import (
	"fmt"
	"net/http"
)

// 有更深一层任务的时候，通过创建
// context.Context withDeadline()
// 来实现超时取消任务

func notFoundHandler(w http.ResponseWriter, r *http.Request) {

}

func middleWareHandlerSample(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

//var fileSystemHandler handlerFunc =
// 	http.FileServer(http.Dir(".")).ServeHTTP

func testRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "RequestURI: %q\n", r.RequestURI)
	fmt.Fprintf(w, "URL: %#v\n", r.URL)
	fmt.Fprintf(w, "Body.ContentLength: %d (-1 means unknown)\n", r.ContentLength)
	fmt.Fprintf(w, "Close: %v (relevant for HTTP/1 only)\n", r.Close)
	fmt.Fprintf(w, "TLS: %#v\n", r.TLS)
	fmt.Fprintf(w, "\nHeaders:\n")
	_ = r.Header.Write(w)
}

func testIndex(w http.ResponseWriter, r *http.Request) {
	_, _ =fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>HTTP/2 Server Push</title>
  <link rel="stylesheet" href="/app.css">
  <script src="/app.js"></script>
</head>
<body>
  <img class="echo" src="/echo.png">
  <h2>The following static files are served via HTTP/2 server push</h2>
  <ul>
    <li><code>/app.css</code></li>
    <li><code>/app.js</code></li>
    <li><code>/echo.png</code></li>
  </ul>
</body>
</html>
`   )
	
	_ = r.Body.Close()
}

func checkUserLoginStatus(w http.ResponseWriter, r *http.Request) {

}

func updateUserStatusToken(w http.ResponseWriter, r *http.Request) {

}