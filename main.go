package main

import (
	"fmt"
	server "github.com/movarr/movarr_server/httpserver"
)

func main() {
	server.StartServer()
	fmt.Println("Hello, World!")
}
