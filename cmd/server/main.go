package main

import (
	"fmt"
	"log"
)

//Run - responsible for the instantiation
//and startup of the go application
func Run() error {
	log.Println("Starting up the application...")
	return nil
}

func main() {
	fmt.Println("Comments Rest Api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
