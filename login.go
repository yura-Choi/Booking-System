package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	gomail "gopkg.in/gomail.v2"
)

var isAdminCookie http.Cookie //사용자에게 보낼 관리자여부
var sessionCookie http.Cookie //사용자에게 보낼 세션키
var codeCookie http.Cookie    //로그인제한해제를 위해 인증번호 발급 시 해당 인증번호 저장을 위한 쿠키
var currentId string

var sessionKeyMap map[string]string = make(map[string]string)

func login(w http.ResponseWriter, req *http.Request) {
	isUserAdmin := req.PostFormValue("isAdmin") //사용자가 관리자인가
	currentId = req.PostFormValue("id")         //로그인을 하려는 사용자가 입력한 아이디
	password := req.PostFormValue("password")   //로그인을 하려는 사용자가 입력한 비밀번호

	//데이터베이스에 저장된 비밀번호와 비교하기 위해 입력받은 비밀번호 해싱
	shaPassword := sha256.New()
	io.WriteString(shaPassword, password)
	shaHex := hex.EncodeToString(shaPassword.Sum(nil))

	//데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	//입력한 id가 관리자테이블에 존재하는지 검사_ 존재한다면 관리자여부를 관리자로 저장
	var comparePassword string
	var q string
	if isUserAdmin == "admin" {
		q = "SELECT password FROM AdminInfo WHERE id=?"

	} else {
		q = "SELECT password FROM MemberInfo WHERE id=?"
	}
	rows, err := db.Query(q, currentId)
	printErr(err)
	defer rows.Close()
	//로그인을 하려는 사용자가 입력한 아이디가 존재할 경우 해당 아이디와 같은 row에 있는 비밀번호를 가져와 comparePassword에 저장
	for rows.Next() {
		err = rows.Scan(&comparePassword)
		printErr(err)
	}

	if comparePassword == "" { //로그인을 하려는 사용자가 입력한 아이디에 대응하는 비밀번호가 없다 = 아이디가 존재하지 않는다.
		w.Write([]byte("WrongId"))
		return
	}

	//여기서부터 해당 아이디의 계정이 존재하는 경우
	var loginSuccessCount int8

	//현재 로그인을 시도하려는 id의 로그인 제한 여부를 데이터베이스에서 받아온다.
	if isUserAdmin == "admin" { //아이디가 관리자테이블에 존재할 경우
		q = "SELECT loginCount FROM AdminInfo WHERE id=?"
	} else { //아이디가 회원테이블에 존재할 경우
		q = "SELECT loginCount FROM MemberInfo WHERE id=?"
	}
	rows, err = db.Query(q, currentId)
	printErr(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&loginSuccessCount)
		printErr(err)
	}

	expiration := time.Now().Add(24 * time.Hour)
	isAdminCookie = http.Cookie{Name: "isAdminCookie", Value: isUserAdmin, Expires: expiration}
	http.SetCookie(w, &isAdminCookie)

	//해당 계정의 로그인이 제한되어있거나 로그인 시도 횟수가 5회를 초과할 경우
	if loginSuccessCount > 5 {
		printErr(err)
		w.Write([]byte("loginRestriction"))
		return
	}

	//비밀번호가 일치하지 않을 때
	if comparePassword != shaHex {
		w.Write([]byte("WrongPassword"))
		loginSuccessCount++
		log.Println(loginSuccessCount)
		if isUserAdmin == "admin" {
			_, err = db.Exec("UPDATE AdminInfo SET loginCount=? WHERE id=?", loginSuccessCount, currentId)
		} else {
			_, err = db.Exec("UPDATE MemberInfo SET loginCount=? WHERE id=?", loginSuccessCount, currentId)
		}
		printErr(err)
		return
	}

	//아이디가 존재하고 비밀번호도 일치할 때_ 로그인 성공
	//8자리 정수로 구성된 세션키 랜덤 생성
	rand.Seed(int64(time.Now().Nanosecond()))
	sessionKey := strconv.Itoa(rand.Intn(100000000) + 10000000)

	shaPassword = sha256.New()
	io.WriteString(shaPassword, sessionKey)
	shaHexSession := hex.EncodeToString(shaPassword.Sum(nil))

	sessionKeyMap[shaHexSession] = currentId

	expiration = time.Now().Add(24 * time.Hour) //세션키와 관리자여부를 쿠키에 저장할 시간 추출
	sessionCookie = http.Cookie{Name: "sessionCookie", Value: shaHexSession, Expires: expiration}
	http.SetCookie(w, &sessionCookie)
	w.Write([]byte("success"))

	if isUserAdmin == "admin" {
		_, err = db.Exec("UPDATE AdminInfo SET loginCount=? WHERE id=?", 0, currentId)
	} else {
		_, err = db.Exec("UPDATE MemberInfo SET loginCount=? WHERE id=?", 0, currentId)
	}
	printErr(err)
}

//로그인_인증코드 발급받고 해당 인증코드 이메일로 전송하는 버튼
func loginGetCode(w http.ResponseWriter, req *http.Request) {
	//7자리의 인증코드 랜덤 생성
	rand.Seed(int64(time.Now().Nanosecond()))
	certification := strconv.Itoa(rand.Intn(10000000) + 1000000)
	log.Println(certification)

	shaCode := sha256.New()
	io.WriteString(shaCode, certification)
	shaHexCode := hex.EncodeToString(shaCode.Sum(nil))

	expiration := time.Now().Add(24 * time.Hour) //인증코드를 쿠키에 저장할 시간 추출
	codeCookie = http.Cookie{Name: "certificationCode", Value: shaHexCode, Expires: expiration}
	http.SetCookie(w, &codeCookie)

	//데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	var email, name string
	var q string

	if http.Cookie(isAdminCookie).Value == "admin" {
		q = "SELECT email, name FROM AdminInfo WHERE id=?"
	} else {
		q = "SELECT email, name FROM MemberInfo WHERE id=?"
	}
	rows, err := db.Query(q, currentId)
	printErr(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&email, &name)
		printErr(err)
	}

	m := gomail.NewMessage()
	m.SetAddressHeader("From", "dbfk1207@gmail.com", "Bus Reservation")
	m.SetAddressHeader("To", email, name)
	m.SetHeader("Subject", "[셔틀버스예약시스템]이메일 인증을 해주세요")
	m.SetBody("text/plain", "인증코드: "+certification)

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, "dbfk1207@gmail.com", "01092874006")

	if err := d.DialAndSend(m); err != nil {
		w.Write([]byte("fail"))
		printErr(err)
		return
	}
	w.Write([]byte("success"))
}

//로그인_사용자가 입력한 인증코드가 일치하는지 확인하는 버튼
func loginSetCode(w http.ResponseWriter, req *http.Request) {
	input := req.PostFormValue("code")

	shaCode := sha256.New()
	io.WriteString(shaCode, input)
	shaHexCode := hex.EncodeToString(shaCode.Sum(nil))
	if shaHexCode != http.Cookie(codeCookie).Value {
		w.Write([]byte("WrongCertificationCode"))
		return
	}
	//데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	if http.Cookie(isAdminCookie).Value == "admin" {
		_, err = db.Exec("UPDATE AdminInfo SET loginCount=? WHERE id=?", 0, currentId)
	} else {
		_, err = db.Exec("UPDATE MemberInfo SET loginCount=? WHERE id=?", 0, currentId)
	}
	printErr(err)

	w.Write([]byte("success"))

}

//아이디찾기
func findID(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	email := req.FormValue("email")
	isAdmin := req.FormValue("isAdmin")
	var dbID string

	//데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	//관리자테이블에서 이름과 이메일이 모두 일치하는 id 찾기
	var q string
	if isAdmin == "admin" {
		q = "SELECT id FROM AdminInfo WHERE name=? AND email=?"
	} else if isAdmin == "member" {
		q = "SELECT id FROM MemberInfo WHERE name=? AND email=?"
	}
	rows, err := db.Query(q, name, email)
	printErr(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&dbID)
		printErr(err)
	}
	if dbID == "" {
		w.WriteHeader(400)
		return
	}
	w.Write([]byte(dbID))
}

var currentUserType string
var currentUserPassword string

//비밀번호찾기
func findPassword(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	id := req.FormValue("id")
	email := req.FormValue("email")
	isAdmin := req.FormValue("isAdmin")
	var dbPassword string
	currentUserType = isAdmin
	currentId = id

	//데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	//관리자테이블에서 이름과 이메일이 모두 일치하는 password 찾기
	var q string
	if isAdmin == "admin" {
		q = "SELECT password FROM AdminInfo WHERE name=? AND id=? AND email=?"
	} else if isAdmin == "member" {
		q = "SELECT password FROM MemberInfo WHERE name=? AND id=? AND email=?"
	}
	rows, err := db.Query(q, name, id, email)
	printErr(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&dbPassword)
		printErr(err)
	}

	if dbPassword == "" {
		w.WriteHeader(400)
		return
	}

	currentUserPassword = dbPassword
	w.WriteHeader(200)
}

//비밀번호 찾기 시 새로운 비밀번호 설정
func setNewPassword(w http.ResponseWriter, req *http.Request) {
	password := req.PostFormValue("password")
	passwordAgain := req.PostFormValue("passwordAgain")

	//데이터베이스 오픈
	db, err := sql.Open("mysql", "root:asdf@tcp(127.0.0.1:3306)/bus")
	printErr(err)
	defer db.Close()

	if password == "" || len(password) < 8 || len(password) > 15 || checkChar(password) {
		w.Write([]byte("WrongPassword"))
		return
	} else if password != passwordAgain {
		w.Write([]byte("Mismatched"))
		return
	} else if password == currentId {
		w.Write([]byte("SameWithId"))
		return
	}

	shaCode := sha256.New()
	io.WriteString(shaCode, password)
	shaHexCode := hex.EncodeToString(shaCode.Sum(nil))
	if shaHexCode == currentUserPassword {
		w.Write([]byte("MatchedPrevPassword"))
		return
	}

	if currentUserType == "admin" {
		_, err = db.Exec("UPDATE AdminInfo SET password=? WHERE password=?", shaHexCode, currentUserPassword)
	} else {
		_, err = db.Exec("UPDATE MemberInfo SET password=? WHERE password=?", shaHexCode, currentUserPassword)
	}
	printErr(err)
	w.Write([]byte("ChangeSuccess"))

}

func printErr(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}
