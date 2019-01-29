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

func listen(topic string) {
	//client := connect("sub", uri)
	client := connect("tcp://localhost:1883", "go-sub-client")
	fmt.Println("Subrcribe on topic '" + topic +")..." )
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}

func main() {

	go listen(topic)

}
