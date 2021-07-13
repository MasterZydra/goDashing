package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var debugmode bool

func main() {
	debug := flag.Bool("debugmode", false, "Debug mode for extended informations")
	port := flag.Int("port", 8080, "Port the server is listening on")
	flag.Parse()
	debugmode = *debug
	// Port
	// - flag
	portStr := fmt.Sprintf("%v", *port)
	// - env
	portEnv := os.Getenv("PORT")
 	if portEnv != "" {
		portStr = portEnv
 	}

	var webroot string
	if os.Getenv("WEBROOT") != "" {
		webroot = filepath.Clean(os.Getenv("WEBROOT")) + string(filepath.Separator)
	} else {
		webroot, _ = os.Getwd()
		webroot = webroot + string(filepath.Separator)
	}

	dash := NewDashing(webroot, portStr, os.Getenv("TOKEN")).Start()
	log.Println("listening on :" + portStr)

	http.Handle("/", dash)

	log.Fatal(http.ListenAndServe(":"+portStr, nil))
}
