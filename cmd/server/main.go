package main

import (
	"os"
	server "server/internal"
)

func main() {
	address := os.Args[1]
	port := os.Args[2]
	directory := os.Args[3]
	server.Start(address, port, directory)
}
