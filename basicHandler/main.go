// main.go basic hanlder
package main

import (
	"io"
	"log"
	"net/http"
)

func myServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", myServer)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
