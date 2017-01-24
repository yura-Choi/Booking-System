package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	http.Handle("/", http.FileServer(http.Dir("/Users/apple/project/src/reservation/resource"))) // index.html
	http.HandleFunc("/asdf", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("asdfasdfasdf"))
	})
	http.HandleFunc("/admincheck/value", adminBtn)                      //회원가입 시 관리자여부 확인
	http.HandleFunc("/join/submit", joinSubmit)                         //회원가입 정보 validate 및 DB 저장
	http.HandleFunc("/login", login)                                    //로그인
	http.HandleFunc("/login/findid", findID)                            //아이디찾기
	http.HandleFunc("/login/findpassword", findPassword)                //비밀번호찾기
	http.HandleFunc("/login/setpassword", setNewPassword)               //새롭게 설정할 비밀번호를 입력받고 설정하기
	http.HandleFunc("/index/logout", logOutPrint)                       //로그아웃
	http.HandleFunc("/index/getCookie", getCookie)                      //현재 쿠키의 값 가져오기
	http.HandleFunc("/login/sendcode", loginGetCode)                    //로그인제한 해제를 위한 인증번호 발급
	http.HandleFunc("/login/checkcode", loginSetCode)                   //로그인제한 해제를 위한 인증번호의 일치여부 검사
	http.HandleFunc("/service/getCurrentStatus", getCurrentAdminCookie) //현재 로그인 한 회원이 관리자인지 회원인지 구분
	http.HandleFunc("/reservation/search/round", searchRoundMember)     //현재 예약가능한 왕복 버스 검색
	http.HandleFunc("/reservation/search/one", searchOneMember)         //현재 예약가능한 편도 버스 검색
	http.HandleFunc("/mypage/getinfo", modifyInfo)                      //회원정보수정을 위해 기존 회원정보 불러오기
	http.HandleFunc("/mypage/modify", modifySubmit)                     //회원정보수정 버튼
	http.HandleFunc("/mypage/withdrawal", withDrawalSubmit)             //회원탈퇴
	http.HandleFunc("/join/checkid", checkId)					        //아이디 중복체크
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// /D:/project/src/reservation/resource
// /Users/imac/work/src/reservation/resource
// /Users/apple/project/src/reservation/resource