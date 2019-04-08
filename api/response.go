package main

import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter) {
	io.WriteString(w, "sendErrorResponse")
}

func sendNormalResponse(w http.ResponseWriter) {
	
}