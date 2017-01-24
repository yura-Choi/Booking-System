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

var (
	startDate string,
	startPlace string,
	startTime string,
	endDate string,
	endPlace string,
	endTime string
)

func keepDataRound(w http.ResponseWriter, req *http.Request){
	id := sessionKeyMap[http.Cookie(sessionCookie).Value]
	type := req.FormValue("type")

	if type == "round" {
		startDate = req.FormValue("startDate");
		startPlace = req.FormValue("startPlace");
		startTime = req.FormValue("startTime");
		endDate = req.FormValue("endDate");
		endPlace = req.FormValue("endPlace");
		endTime = req.FormValue("endTime");


	} else if type == "resort" {
		startDate = req.FormValue("startDate");
		startPlace = req.FormValue("startPlace");
		startTime = req.FormValue("startTime");

	} else if type == "seoul" {
		endDate = req.FormValue("endDate");
		endPlace = req.FormValue("endPlace");
		endTime = req.FormValue("endTime");

	}
}

func keepDataOneResort(w http.ResponseWriter, req *http.Request){
	id := sessionKeyMap[http.Cookie(sessionCookie).Value]
	type := "resort"
}