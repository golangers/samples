// Use of this source code is governed by a GPLv3
// license that can be found in the LICENSE file.

// Compile The Go Source By play.golang.org.
package controllers

import (
	"encoding/json"
	"fmt"
	"golanger.com/framework/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const (
	PLAY    = "http://play.golang.org/p/%s"
	COMPILE = "http://play.golang.org/compile"
	FMT     = "http://play.golang.org/fmt"
	SHARE   = "http://play.golang.org/share"
)

type PagePlay struct {
	Application
}

type Compiled struct {
	Compile_errors string
	Output         string
}

func init() {
	App.RegisterController("play/", PagePlay{})
}

func (p *PagePlay) Init() {
	p.Application.Init()
}

func getCode(html []byte) []byte {
	beginRegStr := `(?Usmi)<textarea(.*)>(.*)</textarea>`

	beginRegex, _ := regexp.Compile(beginRegStr)

	pos := beginRegex.FindAllSubmatch(html, 1024)
	if pos == nil {
		fmt.Println("Nothing matched!")
		return nil
	}

	return pos[0][2]
}

func (p *PagePlay) Index() {
	if p.R.Method != "GET" {
		return
	}
	id := p.GET["p"]
	if id == "" {
		return
	}
	play := fmt.Sprintf(PLAY, id)

	resp, err := http.Get(play)
	if err != nil {
		fmt.Println("Err: ", err)
		return
	}
	defer resp.Body.Close()

	buf, _ := ioutil.ReadAll(resp.Body)
	body := utils.M{}
	body["code"] = strings.TrimSpace(string(getCode(buf)))
	if body["code"] == "" {
		return
	}
	p.Body = body
}

func (p *PagePlay) Compile() {
	if p.R.Method != "POST" {
		return
	}

	data := url.Values{"body": {strings.TrimSpace(p.POST["body"])}}
	resp, err := http.PostForm(COMPILE, data)

	if err != nil {
		m := Compiled{"Error communicating with remote server.", ""}
		ret, _ := json.Marshal(m)
		p.RW.Write(ret)
		return
	}
	defer resp.Body.Close()

	buf, _ := ioutil.ReadAll(resp.Body)
	p.RW.Write(buf)
}

func (p *PagePlay) Fmt() {
	if p.R.Method != "POST" {
		return
	}

	data := url.Values{"body": {strings.TrimSpace(p.POST["body"])}}
	resp, err := http.PostForm(FMT, data)

	if err != nil {
		m := Compiled{"Error communicating with remote server.", ""}
		ret, _ := json.Marshal(m)
		p.RW.Write(ret)
		return
	}
	defer resp.Body.Close()

	buf, _ := ioutil.ReadAll(resp.Body)
	p.RW.Write(buf)
}

func (p *PagePlay) Share() {
	if p.R.Method != "POST" {
		return
	}

	var data string
	for key, _ := range p.R.Form {
		data = key
		break
	}
	resp, err := http.Post(SHARE, p.R.Header.Get("Content-type"), strings.NewReader(data))

	if err != nil {
		m := Compiled{"Error communicating with remote server.", ""}
		ret, _ := json.Marshal(m)
		p.RW.Write(ret)
		return
	}
	defer resp.Body.Close()

	buf, _ := ioutil.ReadAll(resp.Body)
	p.RW.Write(buf)
}
