package main

import (
	. "controllers"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var (
	addr      = flag.String("addr", ":80", "Server port")
	configDir = flag.String("config", "./config", "Directory of config")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()*2 - 1)

	flag.Parse()
	os.Chdir(filepath.Dir(os.Args[0]))
	fmt.Println("Listen server address: " + *addr)
	fmt.Println("Read configuration directory success, directory: " + filepath.Join(filepath.Dir(os.Args[0]), *configDir))

	App.Load(*configDir)
	App.ListenAndServe(*addr, App)
}
