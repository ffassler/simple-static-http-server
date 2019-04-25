package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const DefaultPort = 8080

const (
	CurrentDirectoryErrorExitCode = 1
	ListenAndServeErrorExitCode   = 2
)

func main() {

	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannont get current directory : ", err)
		os.Exit(CurrentDirectoryErrorExitCode)
	}

	directory := flag.String("d", currentDirectory, "Directory which contains static files to be served")
	port := flag.Int("p", DefaultPort, "Port")

	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, GetServeFilename(*directory, r.URL.Path))
	})

	errListenAndServe := http.ListenAndServe(fmt.Sprint(":", *port), nil)

	if errListenAndServe != nil {
		log.Fatal("Cannot start the http server : ", errListenAndServe)
		os.Exit(ListenAndServeErrorExitCode)
	}
}

func GetServeFilename(directory string, path string) string {
	if len(path) > 0 {
		path = path[1:]
	}

	return fmt.Sprint(directory, "/", path)
}
