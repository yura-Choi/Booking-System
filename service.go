package main

import (
	"database/sql"
	"net/http"
)

func getCurrentAdminCookie(w http.ResponseWriter, req *http.Request) {
	currentStatus := http.Cookie(isAdminCookie).Value
	w.Write([]byte(currentStatus))
}

func getCurrentAdminAdmit(w http.ResponseWriter, req *http.Request) {
	currentId := sessionKeyMap[http.Cookie(sessionCookie).Value]

	//데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	var q string
	var checkAdmin string

	q = "SELECT checkAdmin FROM AdminInfo WHERE id=?"
	rows, err := db.Query(q, currentId)
	printErr(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&checkAdmin)
		printErr(err)
	}

	w.Write([]byte(checkAdmin))
}
