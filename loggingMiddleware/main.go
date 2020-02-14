// muxRouter using mux
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// ArticlesCategoryHandler is a function handler
func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	// mux.Vars return all path paramters as map
	log.Print("handle ArticlesCategoryHandler...")
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticlesCategoryHandler)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	srv := &http.Server{
		Handler:      handlers.CompressHandler(loggedRouter),
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
