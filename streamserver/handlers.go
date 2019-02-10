package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func helloHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "<h1>Hello World!</h1>")
}

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, err := template.ParseFiles("./html/upload.html")
	if err != nil {
		log.Println(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid

	video, err := os.Open(vl)
	if err != nil {
		log.Fatal(err)
		sendErrorResponse(w, http.StatusInternalServerError, "InternalServerError")
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, vid, time.Now(), video)

	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is StatusBadRequest")
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Form file error : %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "StatusInternalServerError")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error : %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "StatusInternalServerError")
	}

	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fn, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "StatusInternalServerError")
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}
