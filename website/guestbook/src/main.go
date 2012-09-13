package main

import (
	. "controllers"
	"flag"
	"fmt"
	. "golanger.com/framework/middleware"
	"helper"
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

	if mongoDns, ok := App.Database["MongoDB"]; ok && mongoDns != "" {
		mgoServer := helper.NewMongo(mongoDns)
		defer mgoServer.Close()
		Middleware.Add("db", mgoServer)
	}

	App.AddHeader("Content-Type", "text/html; charset=utf-8")
	App.HandleFavicon()
	App.HandleStatic()
	App.ListenAndServe(*addr, App)
}
