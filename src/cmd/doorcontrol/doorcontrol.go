package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	
	"internal/commons"
	mqtt "github.com/eclipse/paho.mqtt.golang"

)

const clientId  string = "door-control-client"
const topic     string = "devfest/bdm/door/command"

func userInput(client mqtt.Client ) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\nInput command ('O' or 'C') -> ")
		command, _ := reader.ReadString('\n')
		if strings.HasPrefix(command, "O") || strings.HasPrefix(command, "o") {
			fmt.Println("Opening the door...")
			publish(client, "O")
		} else if strings.HasPrefix(command, "C") || strings.HasPrefix(command, "c") {
			fmt.Println("Closing the door...")
			publish(client, "C")
		} else {
			fmt.Println("Invalid command.")
		}
		fmt.Println("")
	}
}
func publish(client mqtt.Client, command string) {
	fmt.Println("Publishing command '" + command + "' (topic = " + topic + ")" ) 
	client.Publish(topic, 0, false, command)
}

func main() {
	fmt.Println("Starting door control..." ) 

	client := commons.Connect(clientId)
	fmt.Println("Connected." ) 
	
	userInput(client)
	
	fmt.Println("End." ) 
}
