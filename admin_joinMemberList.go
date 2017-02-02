package main

import (
	"net/http"
	"database/sql"
	"encoding/json"
)

type joinMemberInfo struct {
	JoinDate string
	Name string
	Id string
	Email string
	Phone string
	Birth string
}

func joinedMemberList(w http.ResponseWriter, req *http.Request){
	// 데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1)/bus")
	printErr(err)
	defer db.Close()

	var countColumn int
	err = db.QueryRow("SELECT COUNT(*) FROM MemberInfo").Scan(&countColumn)
	printErr(err)
	infoArr := make([]joinMemberInfo, countColumn)

	q := "SELECT name, id, email, phone, birth, joindate FROM MemberInfo"
	rows, err := db.Query(q)
	printErr(err)
	defer rows.Close()
	for i := 0; rows.Next(); i++{
		err = rows.Scan(&infoArr[i].Name, &infoArr[i].Id, &infoArr[i].Email, &infoArr[i].Phone, &infoArr[i].Birth, &infoArr[i].JoinDate)
		printErr(err)
	}

	jsonData, err := json.Marshal(infoArr)
	printErr(err)

	w.Header().Set("Content-Type", "application/json; utf-8;")
	w.Write(jsonData)
}