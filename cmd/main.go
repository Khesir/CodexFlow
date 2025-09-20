package main

import (
	server "app/internal"
	"embed"
	"log"
)

var embeddedFrontend embed.FS

// var for all necessary internal values
// Create a app struct
func init() {
	// nessary inits
	// connectt to db
	// static assets file system
	// checker for db schema
	// upgrades?
	// 
}
func main() {
	if err := server.Start(":8080", embeddedFrontend); err != nil {
		log.Println("server stopped:", err)
	}
}
