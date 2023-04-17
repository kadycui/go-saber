package utils

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func middlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 执行handel之前的逻辑
		next.ServeHTTP(w, r)
		// 执行完毕handel的逻辑
	})

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Success!!")

}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Register Success!!")

}

func timeHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", request.Method, request.URL.Path)
		next.ServeHTTP(writer, request)
		log.Printf("Completed %s in %v", request.URL.Path, time.Since(start))
	})

}

func MiddServer() {
	// 创建一个路由转发
	mux := http.NewServeMux()

	// 函数 注册中间件
	mux.Handle("/login", timeHandler(http.HandlerFunc(login)))
	mux.Handle("/register", timeHandler(http.HandlerFunc(register)))

	http.ListenAndServe(":8080", mux)

}
