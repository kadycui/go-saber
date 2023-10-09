package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func printBody(r *http.Response) {
	defer func() { _ = r.Body.Close() }()
	content, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

func requestByParams() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}

	params := make(url.Values)
	params.Add("name", "kadycui")
	params.Add("age", "15")

	fmt.Println(params.Encode())
	request.URL.RawQuery = params.Encode()

	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	printBody(r)

}

func requestByHeard() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("user-agent", "firefox")
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	printBody(r)

}

func main() {
	// 如何设置请求的查询参数 http://httpbin.org/get?name=poloxue&age=18
	// 如何定制请求头 user-agent
	//requestByParams()
	requestByHeard()

}
