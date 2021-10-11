package main

import (
	"IoT_Service/server"
	"log"
)

func main() {
	// Create new server application
	app := server.NewApp()

	// Initialize application
	err := app.Initialize()
	if err != nil {
		log.Print(err)
		app.Stop()
		return
	}

	// Start the services
	err = app.Start()
	if err != nil {
		log.Print(err)
		app.Stop()
	}
}
