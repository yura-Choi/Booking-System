package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type reserveInfo struct {
	Member      string
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
	var obj reserveInfo
	obj.Member = req.FormValue("member")
	obj.BusType = req.FormValue("busType")
	obj.ResortDate.String = req.FormValue("resortDate")
	obj.SeoulDate.String = req.FormValue("seoulDate")
	obj.ResortPlace.String = req.FormValue("resortPlace")
	obj.SeoulPlace.String = req.FormValue("seoulPlace")
	obj.ResortTime.String = req.FormValue("resortTime")
	obj.SeoulTime.String = req.FormValue("seoulTime")

	// 데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	if obj.BusType == "왕복" {
		_, err = db.Exec("DELETE FROM ReservationInfo WHERE id=? AND member=? AND busType=? AND resortDate=? AND seoulDate=? AND resortPlace=? AND seoulPlace=? AND resortTime=? AND seoulTime=?", sessionKeyMap[http.Cookie(sessionCookie).Value], obj.Member, obj.BusType, obj.ResortDate, obj.SeoulDate, obj.ResortPlace, obj.SeoulPlace, obj.ResortTime, obj.SeoulTime)
	} else if obj.BusType == "편도(리조트행)" {
		_, err = db.Exec("DELETE FROM ReservationInfo WHERE id=? AND member=? AND busType=? AND resortDate=? AND resortPlace=? AND resortTime=?", sessionKeyMap[http.Cookie(sessionCookie).Value], obj.Member, obj.BusType, obj.ResortDate, obj.ResortPlace, obj.ResortTime)
	} else if obj.BusType == "편도(서울행)" {
		_, err = db.Exec("DELETE FROM ReservationInfo WHERE id=? AND member=? AND busType=? AND seoulDate=? AND seoulPlace=? AND seoulTime=?", sessionKeyMap[http.Cookie(sessionCookie).Value], obj.Member, obj.BusType, obj.SeoulDate, obj.SeoulPlace, obj.SeoulTime)
	}
	printErr(err)

	w.Write([]byte("success"))
}
