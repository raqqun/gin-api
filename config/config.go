package config

import (
	"log"
	"os"
)

// here we define configuration parameters

// Port returns a string with the http port number
func Port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Println(
			"PORT expects int > 1024.",
			"Used 8080 by default.")
		port = "8080"
	}
	return ":" + port
}
