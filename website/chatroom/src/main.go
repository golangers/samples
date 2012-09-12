package main

import (
	"code.google.com/p/go.net/websocket"
	. "controllers"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	_ "templateFunc"
)

var (
	addr       = flag.String("addr", ":80", "Server port")
	configPath = flag.String("config", "./config/site", "site filepath of config")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()*2 + 1)

	flag.Parse()
	os.Chdir(filepath.Dir(os.Args[0]))
	fmt.Println("Listen server address: " + *addr)
	fmt.Println("Read configuration file success, fithpath: " + filepath.Join(filepath.Dir(os.Args[0]), *configPath))

	App.Load(*configPath)

	http.Handle("/chat", websocket.Handler(BuildConnection))
	go InitChatRoom()

	App.AddHeader("Content-Type", "text/html; charset=utf-8")
	App.HandleFavicon()
	App.HandleStatic()
	App.ListenAndServe(*addr, App)
}
