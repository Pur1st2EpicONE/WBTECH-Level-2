package main

import (
	"log"
)

func main() {

	input, err := processInput()
	if err != nil {
		log.Fatal(err)
	}

	sort(input)

}
