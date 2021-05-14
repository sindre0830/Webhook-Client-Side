package main

import (
	"log"
	"main/api"
	"net/http"
	"os"
)

// Main program.
func main() {
	//get port and branch if there isn't a port and set it to 8081
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	//handle webhook input
	http.HandleFunc("/client/v1/input/", api.InputHandler)
	//handle webhook output
	http.HandleFunc("/client/v1/output/", api.OutputHandler)
	//ends program if it can't open port
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
