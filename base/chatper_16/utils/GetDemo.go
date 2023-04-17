package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Demo1() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://www.baidu.com", nil)

	if err != nil {
		fmt.Println(err)
	}

	response, err := client.Do(request)
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(response.StatusCode)
	fmt.Println(string(res))
}

func Demo2() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://www.baidu.com", nil)

	if err != nil {
		fmt.Println(err)
	}

	// 设置request的Header
	request.Header.Set("Accept", "text/html, application/xhtml+xml, application/xml;q=0.9, ＊/＊;q=0.8")
	request.Header.Set("Accept-Charset", "GBK, utf-8;q=0.7, ＊;q=0.3")
	request.Header.Set("Accept-Encoding", "gzip, deflate, sdch")
	request.Header.Set("Accept-Language", "zh-CN, zh;q=0.8")
	request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")

	response, err := client.Do(request)
	fmt.Printf("%#v", response.Request.Header)

}

func Demo3() {
	// http.Get实际上是DefaultClient.Get(url)，Get函数是高度封装的，只有一个参数url
	// // 对于一般的http Request是可以使用，但是不能定制Request
	response, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}
