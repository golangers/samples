package controllers

import (
	"fmt"
	"net/http"
	"time"
)

type PageIndex struct {
}

// _FILTER => "用来过滤要匹配的Method"，然后根据规则来执行
func (p *PageIndex) Filter_Before() []map[string]string {
	return []map[string]string{
		map[string]string{
			"_FILTER": "getIndex",
			"_ALL":    "allow",
			"Index":   "deny",
		},
		map[string]string{
			"_FILTER": "GetTest",
			"GetRW":   "allow",
		},
		map[string]string{
			"_FILTER": "GetTest",
			"_METHOD": "GET",
			"_PARAM":  "id&u",
			"Index":   "allow",
		},
		map[string]string{
			"_FILTER": "GetTest1",
			"GetW":    "allow",
			"Index":   "allow",
		},
	}
}

func (p *PageIndex) Filter_After() []map[string]string {
	return []map[string]string{
		map[string]string{
			"_FILTER": "after",
			"_ALL":    "allow",
			"GetW":    "deny",
		},
	}
}

func (p *PageIndex) Before() {
	fmt.Println("Before_")
}

func (p *PageIndex) After() {
	fmt.Println("After_")
}

func (p *PageIndex) Init() {
	fmt.Println("init")
}

func (p *PageIndex) RouteDefault() {
	fmt.Println("default")
}

func (p *PageIndex) Before_Index() bool {
	fmt.Println("Before_Index")
	return false
}

func (p *PageIndex) Http_GET_Index() {
	fmt.Println("GET")
}

func (p *PageIndex) Http_POST_Index() bool {
	fmt.Println("POST")
	return true
}

func (p *PageIndex) RouteIndex() {
	p.getGGGGG()
	fmt.Println("index")
	//return true
}

func (p *PageIndex) After_Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("After_Index")
	w.Write([]byte("After_Index"))
}

func (p *PageIndex) RouteTo_List() {
	fmt.Println("2222222")
}

func (p *PageIndex) RouteTime(w http.ResponseWriter) {
	w.Write([]byte("time_before"))
	d, _ := time.ParseDuration("10s")
	time.Sleep(d)
	w.Write([]byte("test_after"))
	fmt.Println("time ok")
}

func (p *PageIndex) RouteGetRW(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	w.Write([]byte(r.URL.String()))
}

func (p *PageIndex) RouteGetW(w http.ResponseWriter) {
	w.Write([]byte("getw"))
}

func (p *PageIndex) RouteGetR(r *http.Request) {
	fmt.Println(r.URL)
}

func (p *PageIndex) GetR() {
	fmt.Println("rrrrr")
}

func (p *PageIndex) getGGGGG() {
	fmt.Println("ggggg")
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
