package controllers

import (
	"fmt"
	"net/http"
	"time"
)

type PageIndex struct {
	Application
}

func init() {
	App.RegisterController("index/", PageIndex{})
}

// _FILTER => "用来过滤要匹配的Method"，然后根据规则来执行
func (p *PageIndex) Before_() []map[string]string {
	return []map[string]string{
		map[string]string{
			"_FILTER": "getIndex",
			"_ALL":    "allow",
			"Index":   "deny",
		},
		map[string]string{
			"_FILTER":  "GetTest",
			"TestPage": "allow",
		},
		map[string]string{
			"_FILTER":  "GetTest1",
			"TestPage": "allow",
			"Index":    "allow",
		},
	}
}

func (p *PageIndex) After_() []map[string]string {
	return []map[string]string{
		map[string]string{
			"_FILTER":  "after",
			"_ALL":     "allow",
			"TestPage": "deny",
		},
	}
}

func (p *PageIndex) Before_Index() {
	fmt.Println("Before_Index")
}

func (p *PageIndex) After_Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("After_Index")
}

// route => , index, index/, index/index
func (p *PageIndex) Index() {
	p.RW.Write([]byte("hello world, welcome to access " + p.R.URL.String()))
}

// route => write, index/write
func (p *PageIndex) Write(w http.ResponseWriter) {
	w.Write([]byte("hello world"))
}

func (p *PageIndex) After_Write(w http.ResponseWriter) {
	w.Write([]byte(", welcome to access this path!"))
}

// route => request, index/request
func (p *PageIndex) Request(r *http.Request) {
	fmt.Println(r.URL)
}

// route => write_request, index/write_request
func (p *PageIndex) WriteRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.String()))
}

func (p *PageIndex) Test() {
	p.RW.Write([]byte("test1"))
	d, _ := time.ParseDuration("10s")
	time.Sleep(d)
	p.RW.Write([]byte("test2"))
}

func (p *PageIndex) Filter_getIndex() {
	fmt.Println("GetIndex")
}

func (p *PageIndex) Filter_GetTest(r *http.Request) {
	fmt.Println("GetTest")
	fmt.Println(r.URL)
}

func (p *PageIndex) Filter_GetTest1() {
	fmt.Println("GetTest1")
}

func (p *PageIndex) Filter_after() {
	fmt.Println("Filter_after")
}
