package server

import (
	asciiart "ascii-art-web-stylize/cmd/internal/asciiart"
	"html/template"
	"net/http"
	"strings"
)

// Struct to hold the error data
type ErrorPageData struct {
	Code     string
	ErrorMsg string
}

// Struct to hold the result data
type ResultPageData struct {
	Input  string
	Banner string
	Result string
}

// Function to render the error page
func errHandler(w http.ResponseWriter, r *http.Request, errData *ErrorPageData, status int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	errorTemp, parseErr := template.ParseFiles("templates/error.html")
	if parseErr != nil {
		http.Error(w, "Error page unavailable", http.StatusInternalServerError)
		return
	}

	errorTemp.Execute(w, errData)
}

func renderTemplate(w http.ResponseWriter, filename string, data interface{}) error {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl.Execute(w, data)
}

// Function to render the main page
func MainHandler(w http.ResponseWriter, r *http.Request) {
	// Validating the request path
	if r.URL.Path != "/" {
		errData := ErrorPageData{Code: "404", ErrorMsg: "PAGE NOT FOUND"}
		errHandler(w, r, &errData, http.StatusNotFound)
		return
	}

	// Validating the request method
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		errData := ErrorPageData{Code: "405", ErrorMsg: "METHOD NOT ALLOWED"}
		errHandler(w, r, &errData, http.StatusMethodNotAllowed)
		return
	}

	if err := renderTemplate(w, "index.html", nil); err != nil {
		errData := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &errData, http.StatusInternalServerError)
		return
	}
}

// Function to render the result page
func ResultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errData := ErrorPageData{Code: "405", ErrorMsg: "METHOD NOT ALLOWED"}
		errHandler(w, r, &errData, http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		errData := ErrorPageData{Code: "400", ErrorMsg: "INVALID REQUEST"}
		errHandler(w, r, &errData, http.StatusBadRequest)
		return
	}

	input := r.PostFormValue("input-text")
	inputValidation := strings.ReplaceAll(input, "\r\n", "")

	for _, letter := range inputValidation {
		if letter < 32 || letter > 126 {
			errData := ErrorPageData{Code: "400", ErrorMsg: "INVALID INPUT"}
			errHandler(w, r, &errData, http.StatusBadRequest)
			return
		}
	}

	banner := r.PostFormValue("banner")
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		errData := ErrorPageData{Code: "400", ErrorMsg: "INVALID BANNER"}
		errHandler(w, r, &errData, http.StatusBadRequest)
		return
	}

	ascii, err := asciiart.AsciiArt(input, banner)
	if err != nil {
		errData := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &errData, http.StatusInternalServerError)
		return
	}

	output := ResultPageData{Input: input, Banner: banner, Result: ascii}
	if err := renderTemplate(w, "templates/asciiart.html", output); err != nil {
		errData := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &errData, http.StatusInternalServerError)
	}
}

// NotFoundHandler handles all undefined routes
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	errData := ErrorPageData{Code: "404", ErrorMsg: "PAGE NOT FOUND"}
	errHandler(w, r, &errData, http.StatusNotFound)
}
