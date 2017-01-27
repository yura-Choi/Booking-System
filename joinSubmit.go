package main

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"
	"unicode"

	"crypto/sha256"

	"io"

	_ "github.com/go-sql-driver/mysql"
)

func checkChar(checkstr string) bool {
	var numbercount = 0
	var alphacount = 0
	for i := range checkstr {
		if unicode.IsNumber(rune(checkstr[i])) {
			numbercount++
			continue
		} else if unicode.IsLower(rune(checkstr[i])) == false {
			return true
		}
		alphacount++
	}
	if numbercount > len(checkstr) || alphacount > len(checkstr) {
		return true
	}
	return false
}

func joinSubmit(w http.ResponseWriter, req *http.Request) {
	id := req.PostFormValue("id")
	name := req.PostFormValue("name")
	password := req.PostFormValue("password")
	passwordAgain := req.PostFormValue("passwordAgain")
	email := req.PostFormValue("email")
	phone := req.PostFormValue("phone")
	birth := req.PostFormValue("birth")

	emailValidate := regexp.MustCompile(`^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$`)
	phoneValidate := regexp.MustCompile(`^(?:(010-\d{4})|(01[1|6|7|8|9]-\d{3,4}))-(\d{4})$`)
	birthValidate := regexp.MustCompile(`^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$`)
	if id == "" || len(id) < 8 || len(id) > 10 || checkChar(id) {
		w.Write([]byte("id"))
	} else if name == "" || len(name) >= 20 {
		w.Write([]byte("name"))
	} else if id == password {
		w.Write([]byte("idpassword"))
	} else if password == "" || len(password) < 8 || len(password) > 15 || checkChar(password) {
		w.Write([]byte("password"))
	} else if passwordAgain == "" || password != passwordAgain {
		w.Write([]byte("passwordAgain"))
	} else if email == "" || !emailValidate.MatchString(email) {
		w.Write([]byte("email"))
	} else if phone == "" || !phoneValidate.MatchString(phone) {
		w.Write([]byte("phone"))
	} else if birth != "" && !birthValidate.MatchString(birth) {
		w.Write([]byte("birth"))
	} else if birth == "" {
		birth = "0000-00-00"
	} else {
		db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
		printErr(err)
		defer db.Close()
		shaPassword := sha256.New()
		io.WriteString(shaPassword, password)
		shaHex := hex.EncodeToString(shaPassword.Sum(nil))
		fmt.Println(shaHex)
		currentTime := time.Now().Format("2006-01-02")
		log.Println(currentTime)
		log.Println(isAdmin)
		if isAdmin == true {
			stmt, err := db.Prepare("INSERT INTO AdminInfo(name, id, password, email, phone, birth, joindate) VALUES(?,?,?,?,?,?,?)")
			printErr(err)
			defer stmt.Close()
			_, err = stmt.Exec(name, id, shaHex, email, phone, birth, string(currentTime))
			if err != nil {
				w.Write([]byte("dbfail"))
			} else {
				w.WriteHeader(200)
				return
			}
		} else if isAdmin == false {
			stmt, err := db.Prepare("INSERT INTO MemberInfo(name, id, password, email, phone, birth, joindate) VALUES(?,?,?,?,?,?,?)")
			printErr(err)
			_, err = stmt.Exec(name, id, shaHex, email, phone, birth, string(currentTime))
			if err != nil {
				w.Write([]byte("dbfail"))
				log.Println(err)
			} else {
				w.WriteHeader(200)
				return
			}
		}
	}
}

func checkId(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("inputId")

	// DB Open
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	var q string
	var dbName string
	// 먼저 관리자 회원의 정보부터 검색
	q = "SELECT name FROM AdminInfo WHERE id=?"
	rows, err := db.Query(q, id)
	printErr(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&dbName)
		printErr(err)
	}

	if dbName != "" {
		w.Write([]byte("alreadyExists"))
		return
	}

	q = "SELECT name FROM MemberInfo WHERE id=?"
	rows, err = db.Query(q, id)
	printErr(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&dbName)
		printErr(err)
	}

	if dbName != "" {
		w.Write([]byte("alreadyExists"))
		return
	} else if id == "" {
		w.Write([]byte("empty"))
		return
	}

	w.Write([]byte("success"))
}

func joinToMember(w http.ResponseWriter, req *http.Request){
	isAdmin = false
	w.Write([]byte("success"))
}