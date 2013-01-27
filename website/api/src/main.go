package main

import (
	. "controllers"
	"flag"
	"fmt"
	"runtime"
)

var (
	addr = flag.String("addr", ":80", "Server port")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()*2 - 1)

	flag.Parse()
	fmt.Println("Listen server address: " + *addr)

	//App.LoadData(``)
	App.ListenAndServe(*addr, App)
}
