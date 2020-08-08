package main

import (
	"os"

	gockify "github.com/tomaszmartin/gockify/src"
)

func main() {
	// For now just pass API key as argument
	apiKey := os.Args[1]
	client := gockify.NewClient(apiKey)
	client.GetWorkspaces()
}
