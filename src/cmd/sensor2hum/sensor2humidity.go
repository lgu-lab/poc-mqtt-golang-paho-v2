package main

import (
	"fmt"
	"time"
	
	"internal/commons"
	
)

const clientId  string = "sensor-02"
const topic     string = "devfest/bdm/humidity"
//const waitDuration int = 3

func main() {

	client := commons.Connect(clientId)

	n := 0 ;
	for {
		n++
		
		// le taux d'humidité dans une maison doit se situer entre 50 et 60 %.
		humidity := 50 + ( n % 10 ) 

		//		msg := fmt.Sprintf("My message %d", n)
		msg := fmt.Sprintf("%d", humidity)
		fmt.Println("PUB ["+topic+"] : " + msg)

		// Publish will publish a message with the specified QoS and content to the specified topic.
		// Returns a token to track delivery of the message to the broker
		//
		// Publish(topic string, qos byte, retained bool, payload interface{}) Token
		
		client.Publish(topic, 0, false, msg )
		
        time.Sleep( 3 * time.Second)
    }
}
