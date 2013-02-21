package main

import (
	_ "controllers"
	"flag"
	"fmt"
	"golanger.com/webrouter"
	"net/http"
	"runtime"
)

var (
	addr = flag.String("addr", ":80", "Server port")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()*2 - 1)

	flag.Parse()
	fmt.Println("Listen server address: " + *addr)

	webrouter.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("favicon.ico")
	})

	webrouter.ListenAndServe(*addr, nil)
}
