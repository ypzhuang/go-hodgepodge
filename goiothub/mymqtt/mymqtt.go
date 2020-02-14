package mymqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"	
	"net/url"
	"time"
)

type MyMQTT struct {
	client mqtt.Client 
}

func (m *MyMQTT) Connect(clientId string, uri *url.URL) {
	log.Printf("try to connect mqtt server %s",uri.Host)
	opts := createClientOptions(clientId, uri)
	m.client = mqtt.NewClient(opts)
	token := m.client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	} 
	log.Printf("connected to %s", uri.Host)	
}

func (m *MyMQTT) IsConnected() {
	return 
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetClientID(clientId)
	// opts.SetConnectRetry(true)
	opts.SetAutoReconnect(true)
	// opts.SetConnectRetryInterval(10*time.Second)
	return opts
}

func (m *MyMQTT) Listen(topic string) {	
	token := m.client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
	if err := token.Error(); err != nil {
		log.Fatal(err)
	} 
	log.Print("Listened")
	
}

func (m *MyMQTT) Publish(topic string, payload string) {
	token := m.client.Publish(topic, 0, false, payload)
	if err := token.Error(); err != nil {
		log.Fatal(err)
	} 
	log.Printf("publish message: %s", payload)	
}
