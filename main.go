package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannont get current directory : ", err)
	}

	directory := flag.String("d", currentDirectory, "Directory which contains static files to be served")
	port := flag.Int("p", 8080, "Port")

	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, fmt.Sprint(*directory, "/", r.URL.Path[1:]))
	})

	log.Fatal("Cannot start the http server : ", http.ListenAndServe(fmt.Sprint(":", *port), nil))
}
