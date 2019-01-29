package main

import (
	"fmt"
	"sync"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	
	"internal/commons"
	
)

const topic    string = "devfest/bdm/#"


func main() {

	client := commons.Connect("sensors-log")

	fmt.Println("Subrcribe on topic '" + topic +"'..." )
	
	var wg sync.WaitGroup
    wg.Add(1)

//    if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
//            //wg.Done()
//            fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()) )
//    }); token.Wait() && token.Error() != nil {
//		fmt.Println("Token error : " + fmt.Sprintln(token.Error()) )
//    }

//    token := client.Subscribe(topic, 0, 
//    	func(client mqtt.Client, msg mqtt.Message) {
//            //wg.Done()
//            fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()) )
//	    })

    token := client.Subscribe(topic, 0, messageReceived )
    
    if token.Wait() && token.Error() != nil {
		fmt.Println("Token error : " + fmt.Sprintln(token.Error()) )
    }

	wg.Wait()
}

func messageReceived(client mqtt.Client, msg mqtt.Message) {
    //wg.Done()
    fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()) )
}
