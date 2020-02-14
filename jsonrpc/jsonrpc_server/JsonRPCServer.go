// RPCServer
package main

import (
	jsonparse "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"·
	"github.com/gorilla/rpc/json"
)

// Args rpc arguments
type Args struct {
	Id string
}

type Book struct {
	Id     string `"json:string,omitempty"`
	Name   string `"json:name,omitempty"`
	Author string `"json:author,omitempty"`
}

// JSONServer
type JSONServer struct{}

// GiveBookDetail is a RPC Server Method
func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book
	raw, readerr := ioutil.ReadFile("./book.json")
	if readerr != nil {
		log.Println("error:", readerr)
		os.Exit(1)
	}

	marshalerr := jsonparse.Unmarshal(raw, &books)
	if marshalerr != nil {
		log.Println("error:", marshalerr)
		os.Exit(1)
	}

	for _, book := range books {
		if book.Id == args.Id {
			*reply = book
			break
		}
	}
	return nil
}

func main() {

	// Create a new RPC server
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(JSONServer), "")

	r := mux.NewRouter()

	r.Handle("/rpc", s)

	log.Fatal(http.ListenAndServe(":1234", r))
}

/**
curl -X POST \
		 http://localhost:1234/rpc \
		 -H 'cache-control: no-cache' \
		 -H 'content-type: application/json' \
		 -d '{
		 "method": "JSONServer.GiveBookDetail",
		 "params": [{
		 "Id": "1234"
		 }],
		 "id": "2"
}'

*/
