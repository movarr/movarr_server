package main

import (
	database "github.com/movarr/movarr_server/database"
	server "github.com/movarr/movarr_server/httpserver"
	routes "github.com/movarr/movarr_server/routes"
)

func main() {
	database.InitialiseDatabase()
	routes.HandleRoutes()
	server.StartServer()
}
