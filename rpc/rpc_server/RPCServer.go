// RPCServer
package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

// Args rpc arguments
type Args struct{}

// TimeServer RPC Server
type TimeServer int64

// GiveServerTime is a RPC Server Method
func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	*reply = time.Now().Unix()
	return nil
}

func main() {
	log.Println("try to register rpc server")
	// Create a new RPC server
	timeserver := new(TimeServer)
	// Register RPC server
	rpc.Register(timeserver)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	log.Fatal(http.Serve(l, nil))
}
