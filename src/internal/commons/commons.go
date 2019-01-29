package commons

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const brokerURI  string = "tcp://test.mosquitto.org:1883"

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


func Connect(clientId string) mqtt.Client {
	fmt.Println("Trying to connect (" + brokerURI + ", " + clientId +")..." )
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected.")
	return client
}
