package main

import "net/http"

var ErrQuoteNotFound = apiError{Err: "quote not found", Status: http.StatusNotFound}
var ErrCreatingQuote = apiError{Err: "failed to create quote", Status: http.StatusInternalServerError}
var ErrUpdatingQuote = apiError{Err: "failed to update quote", Status: http.StatusInternalServerError}
var ErrDeletingQuote = apiError{Err: "failed to delete quote", Status: http.StatusInternalServerError}
var ErrInvalidQuote = apiError{Err: "invalid quote", Status: http.StatusBadRequest}
var ErrMissingId = apiError{Err: "missing id", Status: http.StatusBadRequest}
var ErrMethodNotAllowed = apiError{Err: "method not allowed", Status: http.StatusMethodNotAllowed}

type apiError struct {
	Err    string `json:"error"`
	Status int    `json:"status"`
}

func (e apiError) Error() string {
	return e.Err
}
