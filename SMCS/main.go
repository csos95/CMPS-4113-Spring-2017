package main

import (
	"fmt"
	"github.com/csos95/CMPS-4113-Spring-2017/SMCS/server"
)

func main() {
	fmt.Println("Welcome to SMCS.")

	server := server.NewServer("config.json")

	server.Run()

	fmt.Println("Thank you for using SMCS.")
}