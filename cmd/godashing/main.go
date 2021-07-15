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

	printProgramName()

	log.Println("----------------------------------------")
	log.Println("Startup")
	log.Println("----------------------------------------")

	// Extract assets from executable
	ExtractAssets()

	dash := NewDashing(root, portStr, os.Getenv("TOKEN")).Start()
	log.Println("Listen on http://localhost:" + portStr)
	println()
	log.Println("----------------------------------------")
	log.Println("Running")
	log.Println("----------------------------------------")

	http.Handle("/", dash)

	log.Fatal(http.ListenAndServe(":"+portStr, nil))
}

func printProgramName() {
	println()
	println()
	println(" ..|'''.|          '||''|.                  '||       ||")
	println(".|'     '    ...    ||   ||   ....    ....   || ..   ...  .. ...     ... .")
	println("||    .... .|  '|.  ||    || '' .||  ||. '   ||' ||   ||   ||  ||   || ||")
	println("'|.    ||  ||   ||  ||    || .|' ||  . '|..  ||  ||   ||   ||  ||    |''")
	println(" ''|...'|   '|..|' .||...|'  '|..'|' |'..|' .||. ||. .||. .||. ||.  '||||.") 
	println("                                                                   .|....'")
	println()
}