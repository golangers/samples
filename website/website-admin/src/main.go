package main

import (
	. "controllers"
	"flag"
	"fmt"
	. "golanger.com/middleware"
	"helper"
	"os"
	"path/filepath"
	"runtime"
	_ "templateFunc"
)

var (
	addr      = flag.String("addr", ":80", "Server port")
	configDir = flag.String("config", "./config", "Directory of config")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	flag.Parse()
	os.Chdir(filepath.Dir(os.Args[0]))
	fmt.Println("Listen server address: " + *addr)
	fmt.Println("Read configuration directory success, directory: " + filepath.Join(filepath.Dir(os.Args[0]), *configDir))

	App.Load(*configDir)

	if mongoDns, ok := App.Database["MongoDB"]; ok && mongoDns != "" {
		mgoServer := helper.NewMongo(mongoDns)
		defer mgoServer.Close()
		Middleware.Add("db", mgoServer)
	}

	App.AddHeader("Content-Type", "text/html; charset=utf-8")
	App.ListenAndServe(*addr, App)
}
