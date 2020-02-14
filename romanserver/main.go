// main.go export roman number
package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ypzhuang/romanNumerals"
)

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before rquest phase!")
		// Pass control back to the Handler
		handler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after response phase!")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	urlPathElements := strings.Split(r.URL.Path, "/")
	if urlPathElements[1] == "roman_number" {
		number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
		if number == 0 || number > 10 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Not Found"))
		} else {
			fmt.Fprintf(w, "%s", html.EscapeString(romanNumerals.Numerals[number]))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad request"))
	}
}
func main() {
	// http package has methods for dealing with requests
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", middleware(mainLogicHandler))

	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
