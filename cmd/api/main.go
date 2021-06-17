package main

import (
	"fmt"
	"log"
	"nequi.com/poc-services/cmd/api/bootstrap"
)

func main() {
	fmt.Println("Hello, World! Daniela")
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}