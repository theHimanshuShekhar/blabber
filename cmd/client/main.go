package main

import (
	"log"
	"os"

	"github.com/theHimanshuShekhar/blabber/internal/client"
)

func main() {
	if err := client.Start(); err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
}
