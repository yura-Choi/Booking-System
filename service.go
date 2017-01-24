package main

import (
	"net/http"
)

func getCurrentAdminCookie(w http.ResponseWriter, req *http.Request) {
	currentStatus := http.Cookie(isAdminCookie).Value
	w.Write([]byte(currentStatus))
}
