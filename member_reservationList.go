package main

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"log"
)

type reserveInfo struct {
	Member int8
	BusType string
	ResortDate string
	ResortPlace string
	ResortTime string
	SeoulDate string
	SeoulPlace string
	SeoulTime string
}

func showReserveBusList(w http.ResponseWriter, req *http.Request){
	currentId := sessionKeyMap[http.Cookie(sessionCookie).Value]

	// 데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	var q string	
	var countColumn int
	err = db.QueryRow("SELECT COUNT(*) FROM ReservationInfo WHERE id=?", currentId).Scan(&countColumn)
	printErr(err)
	infoArr := make([]reserveInfo, countColumn)

	q = "SELECT member, busType, resortDate, seoulDate, resortPlace, seoulPlace, resortTime, seoulTime FROM MemberInfo WHERE id=?"
	rows, err := db.Query(q, currentId)
	printErr(err)
	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		err = rows.Scan(&infoArr[i].Member, &infoArr[i].BusType, &infoArr[i].ResortDate, &infoArr[i].SeoulDate, &infoArr[i].ResortPlace, &infoArr[i].SeoulPlace, &infoArr[i].ResortTime, &infoArr[i].SeoulTime)
		printErr(err)
	}

	jsonData, err := json.Marshal(infoArr)
	printErr(err)

	log.Println(string(jsonData))

	w.Header().Set("Content-Type", "application/json; utf-8")
	w.Write(jsonData)
}