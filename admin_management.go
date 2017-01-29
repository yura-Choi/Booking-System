package main

import (
	"net/http"
	"database/sql"
	"encoding/json"
)

type adminInfo struct {
	Name string
	Id string
	Email string
	Phone string
	Birth string
	JoinDate string
}

func showAdminList(w http.ResponseWriter, req *http.Request){
	// 데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	var q string
	var countColumn int
	err = db.QueryRow("SELECT COUNT(*) FROM AdminInfo WHERE checkAdmin=?", "N").Scan(&countColumn)
	printErr(err)
	infoArr := make([]adminInfo, countColumn)

	q = "SELECT name, id, email, phone, birth, joindate FROM AdminInfo WHERE checkAdmin=?"
	rows, err := db.Query(q, "N")
	printErr(err)
	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		err = rows.Scan(&infoArr[i].Name, &infoArr[i].Id, &infoArr[i].Email, &infoArr[i].Phone, &infoArr[i].Birth, &infoArr[i].JoinDate)
		printErr(err)
	}

	jsonData, err := json.Marshal(infoArr)
	printErr(err)

	w.Header().Set("Content-Type", "application/json; utf-8")
	w.Write(jsonData)
}

func processAdmitList(w http.ResponseWriter, req *http.Request){
	typeDo := req.FormValue("typeDo")
	reqId := req.FormValue("id")

	// 데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	if typeDo == "admit" {
		_, err = db.Exec("UPDATE AdminInfo SET checkAdmin=? WHERE id=?", "A", reqId)
	} else if typeDo == "refuse" {
		_, err = db.Exec("UPDATE AdminInfo SET checkAdmin=? WHERE id=?", "R", reqId)
	}
	printErr(err)
}