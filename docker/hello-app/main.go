package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "8080"
	server := http.NewServeMux()
	server.HandleFunc("/", handler)

	log.Printf("hello-app-server start on port:%s.", port)
	err := http.ListenAndServe(":"+port, server)
	log.Fatal(err)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handle request: %s", r.URL.Path)

	fmt.Fprintf(w, "Hello, World!\n")

	host, _ := os.Hostname()
	version := os.Getenv("APP_VERSION")
	fmt.Fprintf(w, "Version: %s\n", version)
	fmt.Fprintf(w, "Hostname: %s\n", host)
}
