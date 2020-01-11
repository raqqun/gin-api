package config

import (
    "os"
    "log"
)

// here we define configuration parameters

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
