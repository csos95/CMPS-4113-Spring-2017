package main

import (
	"fmt"
	"github.com/csos95/CMPS-4113-Spring-2017/SMCS/server"
	"log"
)

func main() {
	fmt.Println("Welcome to SMCS.")

	server := server.NewServer("config.json")

	err := server.Run()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Thank you for using SMCS.")
}
