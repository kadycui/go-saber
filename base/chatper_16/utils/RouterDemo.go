package utils

import "net/http"

type helloHandler struct {
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Golang!"))
}

func Demo8() {
	// 只要是实现了Handler接口的类型，都可以作为路由处理器来实现对请求的处理及响应
	http.Handle("/", &helloHandler{})
	http.ListenAndServe(":8080", nil)
}
