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
var root string

func main() {
	debug := flag.Bool("debugmode", false, "Debug mode for extended informations")
	port := flag.Int("port", 8080, "Port the server is listening on")
	webroot := flag.String("webroot", "", "root path for webserver")
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

	// Web root
	// - flag
	root := *webroot
	if root == "" {
		root, _ = os.Getwd()
		root = root + string(filepath.Separator)
	} else {
		root = filepath.Clean(root) + string(filepath.Separator)
	}
	// - env
	webrootEnv := os.Getenv("WEBROOT")
	if webrootEnv != "" {
		root = filepath.Clean(webrootEnv) + string(filepath.Separator)
	}

	dash := NewDashing(root, portStr, os.Getenv("TOKEN")).Start()
	log.Println("listening on :" + portStr)

	http.Handle("/", dash)

	log.Fatal(http.ListenAndServe(":"+portStr, nil))
}
