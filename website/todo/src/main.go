package main

import (
	. "controllers"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	. "golanger.com/middleware"
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

	sqlite, err := sql.Open("sqlite3", "./data/todo.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer sqlite.Close()
	Middleware.Add("db", sqlite)

	App.AddHeader("Content-Type", "text/html; charset=utf-8")
	App.ListenAndServe(*addr, App)
}
