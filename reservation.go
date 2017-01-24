package main

import (
	"log"
	"net/http"
)

func searchRoundMember(w http.ResponseWriter, req *http.Request) {
	startResort := req.PostFormValue("resortDate")
	startSeoul := req.PostFormValue("seoulDate")
	log.Println(startResort)
	log.Println(startSeoul)
}

func searchOneMember(w http.ResponseWriter, req *http.Request) {
	startDate := req.PostFormValue("startDate")
	log.Println(startDate)
}
