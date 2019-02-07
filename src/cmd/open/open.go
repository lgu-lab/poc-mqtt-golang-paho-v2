package main

import (
	"fmt"
	
	"internal/commons"
	//mqtt "github.com/eclipse/paho.mqtt.golang"

)

const clientId  string = "sensor-03-command"
const topic     string = "devfest/bdm/door/command"

func main() {
	fmt.Println("Starting..." ) 

	client := commons.Connect(clientId)
	
	fmt.Println("Publishing... topic = " + topic ) 
	client.Publish(topic, 0, false, "O")

	fmt.Println("End." ) 
}
