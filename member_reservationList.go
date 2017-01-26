package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type reserveInfo struct {
	Member      int8
	BusType     string
	ResortDate  sql.NullString
	ResortPlace sql.NullString
	ResortTime  sql.NullString
	SeoulDate   sql.NullString
	SeoulPlace  sql.NullString
	SeoulTime   sql.NullString
}

func showReserveBusList(w http.ResponseWriter, req *http.Request) {
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

	q = "SELECT member, busType, resortDate, seoulDate, resortPlace, seoulPlace, resortTime, seoulTime FROM ReservationInfo WHERE id=?"
	rows, err := db.Query(q, currentId)
	printErr(err)
	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		err = rows.Scan(&infoArr[i].Member, &infoArr[i].BusType, &infoArr[i].ResortDate, &infoArr[i].SeoulDate, &infoArr[i].ResortPlace, &infoArr[i].SeoulPlace, &infoArr[i].ResortTime, &infoArr[i].SeoulTime)
		if err != nil {
			panic(err)
		}
	}

	jsonData, err := json.Marshal(infoArr)
	printErr(err)

	w.Header().Set("Content-Type", "application/json; utf-8")
	w.Write(jsonData)
}

func deleteReservation(w http.ResponseWriter, req *http.Request) {

	// 데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()
}
