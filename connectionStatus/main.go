package main

import (
	"fmt"
	"log"
	"net/http"
)

const testIp = "127.0.0.1"

func GetConnectionStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Checking connection status is not implemented yet...")
	fmt.Println("Endpoint Hit: Check getConnectionStatus")
}

func handleRequests() {
	http.HandleFunc("/", getConnectionStatus)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	handleRequests()
}
