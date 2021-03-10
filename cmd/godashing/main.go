package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var debugmode bool

func main() {
	debug := flag.Bool("debugmode", false, "Debug mode for extended informations")
	flag.Parse()
	debugmode = *debug

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	var webroot string
	if os.Getenv("WEBROOT") != "" {
		webroot = filepath.Clean(os.Getenv("WEBROOT")) + string(filepath.Separator)
	} else {
		webroot, _ = os.Getwd()
		webroot = webroot + string(filepath.Separator)
	}

	dash := NewDashing(webroot, port, os.Getenv("TOKEN")).Start()
	log.Println("listening on :" + port)

	http.Handle("/", dash)

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
