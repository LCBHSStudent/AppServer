package requestHandle

import (
	"fmt"
	"net/http"
)

// 有更深一层任务的时候，通过创建
// context.Context withDeadline()
// 来实现超时取消任务

func testMain(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "RequestURI: %s\n", r.RequestURI)
	_, _ = fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	_, _ = fmt.Fprintf(w, "TestMain")
}

func CheckUserLoginStatus(w http.ResponseWriter, r *http.Request) {

}

func UpdateUserStatusToken(w http.ResponseWriter, r *http.Request) {

}