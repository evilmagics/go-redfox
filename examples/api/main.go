package main

import (
	"net/http"

	"github.com/evilmagics/go-redfox"
)

type InternalException redfox.Exception[string]

// Assume this is on your internal exception module
// This is just an simple example without using Manager to handle uncached exceptions
var (
	USERNAME_REQUIRED = redfox.NewForAPI("USERNAME_REQUIRED", "username must filled", http.StatusBadRequest)
	PASSWORD_REQUIRED = redfox.NewForAPI("PASSWORD_REQUIRED", "password must filled", http.StatusBadRequest)
	SERVER_ERROR      = redfox.NewForAPI("SERVER_ERROR", "internal server error", http.StatusInternalServerError)
)

func validation() InternalException {
	// Assume username is empty
	return USERNAME_REQUIRED
}

func http_handler(w http.ResponseWriter, r *http.Request) {
	// Validation simulation
	err := validation()

	w.WriteHeader(err.StatusCode())
	w.Write([]byte(err.Message()))
}

func main() {
	http.HandleFunc("/", http_handler)
	http.ListenAndServe(":8080", nil)
}
