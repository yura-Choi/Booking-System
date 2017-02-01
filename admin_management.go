package main

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"gopkg.in/gomail.v2"
	"crypto/sha256"
	"io"
	"encoding/hex"
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
	var email, name string

	// 데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	q := "SELECT email, name FROM AdminInfo WHERE id=?"
	rows, err := db.Query(q, reqId)
	printErr(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&email, &name)
		printErr(err)
	}

	m := gomail.NewMessage()
	m.SetAddressHeader("From", "dbfk1207@gmail.com", "Bus Reservation")
	m.SetAddressHeader("To", email, name)
	m.SetHeader("subject", "[셔틀버스예약시스템]안내메일입니다.")

	if typeDo == "admit" {
		_, err = db.Exec("UPDATE AdminInfo SET checkAdmin=? WHERE id=?", "A", reqId)
		m.SetBody("text/plain", name+" 관리자님의 신청이 승인되었습니다. 이제 관리자메뉴를 이용하실 수 있습니다.")
	} else if typeDo == "refuse" {
		_, err = db.Exec("UPDATE AdminInfo SET checkAdmin=? WHERE id=?", "R", reqId)
		m.SetBody("text/plain", name+" 관리자님의 신청이 거절되었습니다. 정상적인 처리를 원하신다면 관리자 연락처로 문의해주세요(문의: 02-000-0000)")
	}
	d := gomail.NewPlainDialer("smtp.gmail.com", 587, "dbfk1207@gmail.com", "01092874006")

	if err := d.DialAndSend(m); err != nil {
		w.WriteHeader(400)
		printErr(err)
		return
	}
	printErr(err)
}

func checkPasswordAdmin(w http.ResponseWriter, req *http.Request){
	inputPassword := req.FormValue("password")
	currentId := sessionKeyMap[http.Cookie(sessionCookie).Value]
	var password string

	//데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	q := "SELECT password FROM AdminInfo WHERE id=?"
	rows, err := db.Query(q, currentId)
	printErr(err)
	defer rows.Close()
	for rows.Next(){
		err = rows.Scan(&password)
		printErr(err)
	}

	shaPassword := sha256.New()
	io.WriteString(shaPassword, inputPassword)
	shaHex := hex.EncodeToString(shaPassword.Sum(nil))

	if shaHex == password {
		w.Write([]byte("correct"))
	} else {
		w.Write([]byte("incorrect"))
	}
}