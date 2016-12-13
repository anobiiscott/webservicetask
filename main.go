package main

import (
	"flag"
	"net/http"
	"github.com/scottbeaman/webservice-exercise/handlers"
	"log"
)

func main() {
	port := flag.String("port", "8000", "Server port")
	flag.Parse()
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":"+*port, handlers.Router()))
}
