package main

import (
	"log"

	"github.com/Wmengti/go-programming-tool/tour/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
