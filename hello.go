package main

import (
	"fmt"
	"log"

	"github.com/arcinston/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// Get a greeting message and print it.
	messages, err := greetings.Hellos([]string{"arush", "om", "dev"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
