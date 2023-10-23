package main

import "net/http"

var ErrQuoteNotFound = apiError{Err: "quote not found", Status: http.StatusNotFound}
var ErrMethodNotAllowed = apiError{Err: "method not allowed", Status: http.StatusMethodNotAllowed}

type apiError struct {
	Err    string `json:"error"`
	Status int    `json:"status"`
}

func (e apiError) Error() string {
	return e.Err
}
