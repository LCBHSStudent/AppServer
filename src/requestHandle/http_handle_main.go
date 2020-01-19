package requestHandle

import (
	"net"
	"net/http"
)

var _listener []net.Listener
var _http2Mux *http.ServeMux
