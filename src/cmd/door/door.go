package main

import (
	"fmt"
	"sync"
	"strings"
	
	"internal/commons"
	mqtt "github.com/eclipse/paho.mqtt.golang"

)

const clientId  string = "door-client"
const topicSUB  string = "devfest/bdm/door/command"
const OPEN      string = "O"
const CLOSED    string = "C"

var doorState string = OPEN // Initial state is 'OPEN'

func subscribe(client mqtt.Client) {
	fmt.Println("Subscribe on topic '" + topicSUB +")..." )
	
//    if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
//            //wg.Done()
//            fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()) )
//    }); token.Wait() && token.Error() != nil {
//    	//log.Fatal(token.Error())
//            //t.Fatal(token.Error())
//		fmt.Println("Fatal error '" + token.Error().Error() +")..." )
//    }

    token := client.Subscribe(topicSUB, 0, onMessage )
//    	func(client mqtt.Client, msg mqtt.Message) {
//			fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()) )
//			text := string(msg.Payload()) 
//			if strings.HasPrefix(text, "O") || strings.HasPrefix(text, "o"){
//				open()
//			}
//			if strings.HasPrefix(text, "C") || strings.HasPrefix(text, "c"){
//				close()
//			}
//	    })
    
    if token.Wait() && token.Error() != nil {
		fmt.Println("Fatal error '" + token.Error().Error() +")..." )
    }

}

func onMessage(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("")
	fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()) )
	text := string(msg.Payload()) 
	processCommand(text)
}

func processCommand(command string) {
	if strings.HasPrefix(command, "O") || strings.HasPrefix(command, "o") {
		open()
	} else if strings.HasPrefix(command, "C") || strings.HasPrefix(command, "c") {
		close()
	} else {
		fmt.Println("Invalid command.")
	}
}
func open() {
	if doorState != OPEN {
		fmt.Println("Opening the door...")
		doorState = OPEN
		printDoorState()
		publishDoorState(doorState)
	} else {
		fmt.Println("Door is already open.")
	}
}
func close() {
	if doorState != CLOSED {
		fmt.Println("Closing the door...")
		doorState = CLOSED
		printDoorState()
		publishDoorState(doorState)
	} else {
		fmt.Println("Door is already closed.")
	}
}
func printDoorState() {
	state := "OPEN" 
	if doorState != OPEN {
		state = "CLOSED"
	}
	fmt.Println("Door state is '" + state + "'")
}
func publishDoorState(state string) {
	fmt.Println("Publishing new door state : " + state )
}
func main() {
	fmt.Println("Starting..." )

	client := commons.Connect(clientId)

	var wg sync.WaitGroup
    wg.Add(1) // wait group for 2 go routines
	
	subscribe(client)
	
	printDoorState()
	
	fmt.Println("Waiting for commands..." )
	wg.Wait()
}
