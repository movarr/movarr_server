package main

import (
	server "github.com/movarr/movarr_server/httpserver"
	routes "github.com/movarr/movarr_server/routes"
)

func main() {
	routes.HandleRoutes()
	server.StartServer()
}
