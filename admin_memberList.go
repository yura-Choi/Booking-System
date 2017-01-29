package main

import (
	"net/http"
	"database/sql"
	"encoding/json"
)

type reserveMemberInfo struct {
	Time string
	Member string
	Name string
	Phone string
	Birth string
}

func showReservedMemberList(w http.ResponseWriter, req *http.Request){
	busway := req.FormValue("busway")
	date := req.FormValue("date")
	resortBus := req.FormValue("resortBus")
	seoulBus := req.FormValue("seoulBus")
	var id string

	// 데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	var countColumn int
	var infoArr []reserveMemberInfo
	if busway == "resort" {
		err = db.QueryRow("SELECT COUNT(*) FROM ReservationInfo WHERE (busType=? AND resortDate=? AND resortPlace=?) OR (busType=? AND resortDate=? AND resortPlace=?)", "편도(리조트행)", date, resortBus, "왕복", date, resortBus).Scan(&countColumn)
		printErr(err)
		infoArr = make([]reserveMemberInfo, countColumn)
		
		q := "SELECT id, member, resortTime FROM ReservationInfo WHERE (busType=? AND resortDate=? AND resortPlace=?) OR (busType=? AND resortDate=? AND resortPlace=?)"
		rows, err := db.Query(q, "편도(리조트행)", date, resortBus, "왕복", date, resortBus)
		printErr(err)
		defer rows.Close()
		for i := 0; rows.Next(); i++ {
			err = rows.Scan(&id, &infoArr[i].Member, &infoArr[i].Time)
			printErr(err)

			qq := "SELECT name, phone, birth FROM MemberInfo WHERE id=?"
			row, err := db.Query(qq, id)
			printErr(err)
			defer row.Close()
			for row.Next() {
				err = row.Scan(&infoArr[i].Name, &infoArr[i].Phone, &infoArr[i].Birth)
				printErr(err)
			}
		}
	} else if busway == "seoul" {
		err = db.QueryRow("SELECT COUNT(*) FROM ReservationInfo WHERE (busType=? AND seoulDate=? AND seoulPlace=?) OR (busType=? AND seoulDate=? AND seoulPlace=?)", "편도(서울행)", date, seoulBus, "왕복", date, seoulBus).Scan(&countColumn)
		printErr(err)
		infoArr = make([]reserveMemberInfo, countColumn)

		q := "SELECT id, member, seoulTime FROM ReservationInfo WHERE (busType=? AND seoulDate=? AND seoulPlace=?) OR (busType=? AND seoulDate=? AND seoulPlace=?)"
		rows ,err := db.Query(q, "편도(서울행)", date, seoulBus, "왕복", date, seoulBus)
		printErr(err)
		defer rows.Close()
		for i := 0; rows.Next(); i++ {
			err = rows.Scan(&id, &infoArr[i].Member, &infoArr[i].Time)
			printErr(err)

			qq := "SELECT name, phone, birth FROM MemberInfo WHERE id=?"
			row, err := db.Query(qq, id)
			printErr(err)
			defer row.Close()
			for row.Next() {
				err = row.Scan(&infoArr[i].Name, &infoArr[i].Phone, &infoArr[i].Birth)
				printErr(err)
			}
		}
	}

	jsonData, err := json.Marshal(infoArr)
	printErr(err)

	w.Header().Set("Content-Type", "application/json; utf-8")
	w.Write(jsonData)
}