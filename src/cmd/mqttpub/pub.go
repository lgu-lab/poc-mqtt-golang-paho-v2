package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Local constants 
//const brokerURI  string = "tcp://localhost:1883"
//const clientId string = "xxxx"
//const user     string = "xxxx"
//const password string = "xxxx"
const topic    string = "devfest/bdm"

func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()

	// AddBroker adds a broker URI to the list of brokers to be used. 
	// The format should be  "scheme://host:port"
	opts.AddBroker(brokerURI)
	
//	opts.SetUsername(user)
//	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}


func connect(brokerURI string, clientId string) mqtt.Client {
	fmt.Println("Trying to connect (" + brokerURI + ", " + clientId +")..." )
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func main() {

	client := connect("tcp://localhost:1883", "go-pub-client")

//	ticker := time.NewTicker(1 * time.Second)
//	for t := range ticker.C {  // 'C' is a channel that delivers 'ticks' of a clock at intervals
//		msg := t.String()
//		fmt.Println("PUB : " + msg)
//		client.Publish(topic, 0, false, msg)
//	}
	
	n := 0 ;
	for {
		n++
		msg := fmt.Sprintf("My message %d", n)
		fmt.Println("PUB : " + msg)
		
		// Publish will publish a message with the specified QoS and content to the specified topic.
		// Returns a token to track delivery of the message to the broker
		//
		// Publish(topic string, qos byte, retained bool, payload interface{}) Token
		
		client.Publish(topic, 0, false, msg)
		
        time.Sleep(time.Second)
    }
}
