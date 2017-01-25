package main

import "net/http"

func logOutPrint(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(http.Cookie(sessionCookie).Value))
	delete(sessionKeyMap, http.Cookie(sessionCookie).Value)
}

func getCookie(w http.ResponseWriter, req *http.Request) {
	exists := sessionKeyMap[http.Cookie(sessionCookie).Value]
	w.Write([]byte(exists))
}
