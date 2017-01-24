package main

import (
	"log"
	"net/http"
	"os"
)

var isAdmin bool

func adminBtn(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("resource/config.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		log.Println(err)
		return
	}
	var data = make([]byte, fi.Size())
	n, err := file.Read(data)
	if err != nil {
		log.Println(err)
		return
	}
	_ = n

	code := req.FormValue("adminCheck")
	if code == string(data) {
		isAdmin = true
		w.WriteHeader(200)
	} else if code != string(data) {
		isAdmin = false
		w.WriteHeader(400)
	}
}
