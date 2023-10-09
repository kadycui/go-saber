package utils

import (
	"fmt"
	"log"
	"net/http"
)

func myfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello golang")
}

func HttpServer() {
	http.HandleFunc("/hello", myfunc)
	log.Fatal(http.ListenAndServe(":8088", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Hello")

}

func FileServer() {

	http.HandleFunc("/", index)

	// 静态文件路径
	fsh := http.FileServer(http.Dir("D:\\Go\\src\\github.com/kadycui/go-saber\\chatper_16\\static\\"))
	http.Handle("/static/", http.StripPrefix("/static/", fsh))
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
