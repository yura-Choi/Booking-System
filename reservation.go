package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type busReservationData struct {
	BusType    string
	Id         string
	StartDate  string
	StartPlace string
	StartTime  string
	EndDate    string
	EndPlace   string
	EndTime    string
	Member     string
}

var infoBus busReservationData

func saveReservationData(w http.ResponseWriter, req *http.Request) {
	infoBus.Id = sessionKeyMap[http.Cookie(sessionCookie).Value]
	infoBus.BusType = req.FormValue("busType")

	if infoBus.BusType == "round" {
		infoBus.StartDate = req.FormValue("startDate")
		infoBus.StartPlace = req.FormValue("startPlace")
		infoBus.StartTime = req.FormValue("startTime")
		infoBus.EndDate = req.FormValue("endDate")
		infoBus.EndPlace = req.FormValue("endPlace")
		infoBus.EndTime = req.FormValue("endTime")
		w.Write([]byte("success"))
	} else if infoBus.BusType == "resort" {
		infoBus.StartDate = req.FormValue("startDate")
		infoBus.StartPlace = req.FormValue("startPlace")
		infoBus.StartTime = req.FormValue("startTime")
		w.Write([]byte("success"))
	} else if infoBus.BusType == "seoul" {
		infoBus.EndDate = req.FormValue("endDate")
		infoBus.EndPlace = req.FormValue("endPlace")
		infoBus.EndTime = req.FormValue("endTime")
		w.Write([]byte("success"))
	} else {
		w.Write([]byte("error"))
	}

	log.Println(infoBus)
}

func printReservationResult(w http.ResponseWriter, req *http.Request) {
	jsonData, err := json.Marshal(infoBus)
	printErr(err)

	w.Header().Set("Content-Type", "application/json; utf-8")
	w.Write(jsonData)
}

func submitReservationData(w http.ResponseWriter, req *http.Request) {
	infoBus.Member = req.FormValue("member")

	//데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	if infoBus.BusType == "round" {
		stmt, err := db.Prepare("INSERT INTO ReservationInfo(id, member, busType, resortDate, seoulDate, resortPlace, seoulPlace, resortTime, seoulTime) VALUES(?,?,?,?,?,?,?,?,?)")
		printErr(err)
		defer stmt.Close()
		_, err = stmt.Exec(infoBus.Id, infoBus.Member, "왕복", infoBus.StartDate, infoBus.EndDate, infoBus.StartPlace, infoBus.EndPlace, infoBus.StartTime, infoBus.EndTime)
		if err != nil {
			w.WriteHeader(400)
			log.Println(err)
		} else {
			w.WriteHeader(200)
			return
		}
	} else if infoBus.BusType == "resort" {
		stmt, err := db.Prepare("INSERT INTO ReservationInfo(id, member, busType, resortDate, resortPlace, resortTime) VALUES(?,?,?,?,?,?)")
		printErr(err)
		defer stmt.Close()
		_, err = stmt.Exec(infoBus.Id, infoBus.Member, "편도(리조트행)", infoBus.StartDate, infoBus.StartPlace, infoBus.StartTime)
		if err != nil {
			w.WriteHeader(400)
			log.Println(err)
		} else {
			w.WriteHeader(200)
			return
		}
	} else if infoBus.BusType == "seoul" {
		stmt, err := db.Prepare("INSERT INTO ReservationInfo(id, member, busType, seoulDate, seoulPlace, seoulTime) VALUES(?,?,?,?,?,?)")
		printErr(err)
		defer stmt.Close()
		_, err = stmt.Exec(infoBus.Id, infoBus.Member, "편도(서울행)", infoBus.EndDate, infoBus.EndPlace, infoBus.EndTime)
		if err != nil {
			w.WriteHeader(400)
			log.Println(err)
		} else {
			w.WriteHeader(200)
			return
		}
	}

}
