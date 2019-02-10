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

var doorStatus string = "O"

func subscribeAndWait( client mqtt.Client) {
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
	}
	if strings.HasPrefix(command, "C") || strings.HasPrefix(command, "c") {
		close()
	}
}
func open() {
	fmt.Println("Opening the door...")
	doorStatus = "O"
	fmt.Println("Door is OPEN")
	fmt.Println("")
}
func close() {
	fmt.Println("Closing the door...")
	doorStatus = "C"
	fmt.Println("Door is CLOSED")
	fmt.Println("")
}

func main() {
	fmt.Println("Starting..." )

	client := commons.Connect(clientId)

	var wg sync.WaitGroup
    wg.Add(1) // wait group for 2 go routines
	
	subscribeAndWait(client)

	//fmt.Println("after subscribeAndWait()" )
	fmt.Println("Waiting..." )
	wg.Wait()
}
