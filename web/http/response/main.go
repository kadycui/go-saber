package main

import (
	"fmt"
	"io"
	"net/http"
)

func responseBody(r *http.Response) {
	context, _ := io.ReadAll(r.Body)
	fmt.Printf("%s", context)

}

func status(r *http.Response) {
	fmt.Println(r.StatusCode) // 状态码
	fmt.Println(r.Status)     // 状态描述

}

func header(r *http.Response) {

}

func encoding(r *http.Response) {

}

func main() {
	r, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}

	defer func() { _ = r.Body.Close() }()

	status(r)

}
