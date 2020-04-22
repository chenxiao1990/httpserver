package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func logmiddlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		path := r.URL.Path
		raw := r.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}
		next.ServeHTTP(w, r)

		end := time.Now()
		ip := r.RemoteAddr

		method := r.Method

		fmt.Printf("%s  %s  %vms %s \n", ip, method, (end.UnixNano()-start.UnixNano())/1000/1000, path)
	})
}

func main() {
	port := flag.Int("port", 8090, "-port=8090")
	dir := flag.String("rootdir", ".", "-rootdir=./dirname")
	flag.Parse()
	fmt.Println("http服务根目录:", *dir)
	fmt.Printf("访问地址: 127.0.0.1:%d \n", *port)

	http.Handle("/", logmiddlewareHandler(http.FileServer(http.Dir(*dir))))
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		panic(err.Error())
	}
}
