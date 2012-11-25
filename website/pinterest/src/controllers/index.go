package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"golanger.com/framework/utils"
	"io/ioutil"
	. "models"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type PageIndex struct {
	Application
}

func init() {
	App.RegisterController("index/", PageIndex{})
}

func (p *PageIndex) Index(w http.ResponseWriter, r *http.Request) {
	body := utils.M{}
	body["classes"], _ = GetClasses()
	body["images"], _ = GetImagesLists()

	var classId = int64(0)
	sClassId, ok := p.GET["classId"]
	if ok {
		classId, _ = strconv.ParseInt(sClassId, 0, 64)
		body["images"], _ = GetImagesListsWithClassId(classId)
	}
	body["current"] = classId

	p.Body = body
}

func (p *PageIndex) Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		buf := new(bytes.Buffer)
		tnow := time.Now()
		binary.Write(buf, binary.LittleEndian, tnow.UnixNano())
		fileName := strings.TrimRight(base64.URLEncoding.EncodeToString(buf.Bytes()), "=")
		filePath := p.UploadDirectory + "images/"
		os.MkdirAll(filePath, 0777)
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}

		iClass, _ := strconv.Atoi(p.POST["class"])
		class := int64(iClass)

		fileExt := strings.ToLower(path.Ext(fileHeader.Filename))
		fileContent, _ := ioutil.ReadAll(file)
		ioutil.WriteFile(filePath+fileName+fileExt, fileContent, 0777)

		go SaveImages(Images{
			Class:      class,
			Name:       fileName,
			Ext:        fileExt,
			Path:       filePath[len(p.AssetsDirectory)-1:],
			Status:     1,
			CreateTime: time.Now().Unix(),
		})

		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (p *PageIndex) Report(w http.ResponseWriter, r *http.Request) {
	id, ok := p.GET["id"]
	if ok {
		idValue, _ := strconv.Atoi(id)
		image, err := GetImage(idValue)
		if err == nil {
			image.Status = 0
			go SaveImages(*image)
		}
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (p *PageIndex) Page(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if pg, ok := p.GET["page"]; ok {
			ipg, _ := strconv.Atoi(pg)
			body := utils.M{}
			body["images"], _ = GetImagesLists(ipg)
			p.Body = body
		}
	}

}
