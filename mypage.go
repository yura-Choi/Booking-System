package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
)

type information struct {
	Name   string
	Email  string
	Phone  string
	Birth  string
	UserId string
}

func modifyInfo(w http.ResponseWriter, req *http.Request) {
	var info information
	id := sessionKeyMap[http.Cookie(sessionCookie).Value]

	//데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	log.Println(http.Cookie(isAdminCookie).Value)
	var q string
	if http.Cookie(isAdminCookie).Value == "admin" {
		q = "SELECT name, email, phone, birth FROM AdminInfo WHERE id=?"
	} else {
		q = "SELECT name, email, phone, birth FROM MemberInfo WHERE id=?"
	}

	rows, err := db.Query(q, id)
	printErr(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&info.Name, &info.Email, &info.Phone, &info.Birth)
		printErr(err)
	}
	info.UserId = id

	jsonData, err := json.Marshal(info)
	printErr(err)

	w.Header().Set("Content-Type", "application/json; utf-8")
	w.Write(jsonData)
}

func modifySubmit(w http.ResponseWriter, req *http.Request) {
	password := req.FormValue("password")
	passwordAgain := req.FormValue("passwordAgain")
	email := req.FormValue("email")
	phone := req.FormValue("phone")
	birth := req.FormValue("birth")
	id := sessionKeyMap[http.Cookie(sessionCookie).Value]

	//데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	// Validate
	emailValidate := regexp.MustCompile(`^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$`)
	phoneValidate := regexp.MustCompile(`^(?:(010-\d{4})|(01[1|6|7|8|9]-\d{3,4}))-(\d{4})$`)
	birthValidate := regexp.MustCompile(`^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$`)
	if id == password {
		w.Write([]byte("idpassword"))
		return
	} else if password == "" || len(password) < 8 || len(password) > 15 || checkChar(password) {
		w.Write([]byte("password"))
		return
	} else if passwordAgain == "" || password != passwordAgain {
		w.Write([]byte("passwordAgain"))
		return
	} else if email == "" || !emailValidate.MatchString(email) {
		w.Write([]byte("email"))
		return
	} else if phone == "" || !phoneValidate.MatchString(phone) {
		w.Write([]byte("phone"))
		return
	} else if birth != "" && !birthValidate.MatchString(birth) {
		w.Write([]byte("birth"))
		return
	} else if birth == "" {
		birth = ""
	}

	shaPassword := sha256.New()
	io.WriteString(shaPassword, password)
	shaHex := hex.EncodeToString(shaPassword.Sum(nil))

	if http.Cookie(isAdminCookie).Value == "admin" {
		_, err = db.Exec("UPDATE AdminInfo SET password=?, email=?, phone=?, birth=? WHERE id=?", shaHex, email, phone, birth, id)
	} else {
		_, err = db.Exec("UPDATE MemberInfo SET password=?, email=?, phone=?, birth=? WHERE id=?", shaHex, email, phone, birth, id)
	}
	printErr(err)
	w.Write([]byte("success"))
}

func withDrawalSubmit(w http.ResponseWriter, req *http.Request) {
	password := req.FormValue("password")
	currentId = sessionKeyMap[http.Cookie(sessionCookie).Value]

	//데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	var q string
	if http.Cookie(isAdminCookie).Value == "admin" {
		q = "SELECT password FROM AdminInfo WHERE id=?"
	} else {
		q = "SELECT password FROM MemberInfo WHERE id=?"
	}

	var comparePassword string
	rows, err := db.Query(q, currentId)
	printErr(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&comparePassword)
		printErr(err)
	}

	shaPassword := sha256.New()
	io.WriteString(shaPassword, password)
	shaHex := hex.EncodeToString(shaPassword.Sum(nil))

	if shaHex != comparePassword {
		w.Write([]byte("disMatched"))
		return
	}

	delete(sessionKeyMap, http.Cookie(sessionCookie).Value)
	if http.Cookie(isAdminCookie).Value == "admin" {
		_, err = db.Exec("DELETE FROM AdminInfo WHERE id=?", currentId)
	} else {
		_, err = db.Exec("DELETE FROM AdminInfo WHERE id=?", currentId)
	}
	printErr(err)

	w.Write([]byte("success"))
}
