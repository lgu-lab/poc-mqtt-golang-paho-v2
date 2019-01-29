package main

import (
	"fmt"
	"time"
	
	"internal/commons"
	
)

const clientId  string = "sensor-01"
const topic     string = "devfest/bdm/temperature"


func main() {

	client := commons.Connect(clientId)

	n := 0 ;
	for {
		n++
		
		// Temperature : 0 à 30
		temperature := n % 30 

		msg := fmt.Sprintf("%d", temperature)
		fmt.Println("PUB ["+topic+"] : " + msg)
		
		// Publish will publish a message with the specified QoS and content to the specified topic.
		// Returns a token to track delivery of the message to the broker
		//
		// Publish(topic string, qos byte, retained bool, payload interface{}) Token
		
		client.Publish(topic, 0, false, msg)
		
        time.Sleep( 1 * time.Second)
    }
}
