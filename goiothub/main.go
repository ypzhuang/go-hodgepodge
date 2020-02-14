package main

import (
	"fmt"	
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ypzhuang/goiothub/mymqtt"
)


func main() {
	topic := "home/living/humdity"
	broker := "tcp://emqx.bdease.com:1883"


	uri, err := url.Parse(broker)
	if err != nil {
		log.Fatal(err)
	}

	var client mymqtt.MyMQTT
	client.Connect(fmt.Sprintf("goiothub_%s",strings.ToLower(randomString(10))),uri)

	go client.Listen(topic)

    go mockpub(client, topic)
	
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe(":8000", nil))


}

func handler( w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Iothub Go is running\n")
}

func mockpub(client mymqtt.MyMQTT, topic string) {
    for i :=0; i < 10; i++ {		
		payload := fmt.Sprintf("{'h': 78, 't': 30, 's_h': 255,'s_status': 'very wet','d':'%v'}",time.Now().Format("2006-01-02 15:04:05"))
		client.Publish(topic, payload)
		time.Sleep(10 * time.Second)
	}	
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        bytes[i] = byte(randomInt(65, 90))
    }
    return string(bytes)
}