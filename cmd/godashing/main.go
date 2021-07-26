package main

import (
	"flag"
	"fmt"
	"godashing"
	"godashing/internal/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var debugmode bool
var root string

func main() {
	// Flags
	debug := flag.Bool("debugmode", false, "Debug mode for extended informations")
	port := flag.Int("port", 8080, "Port the server is listening on")
	webroot := flag.String("webroot", "", "root path for webserver")
	log2file := flag.Bool("log2file", false, "Save log output into a file")
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

	if *log2file {
		// Create path if it not exists
		if ok, _ := utils.Exists(godashing.LogFileName); !ok {
			utils.Create(godashing.LogFileName)
		}

		// Write log to file instead of console window
		// Source: https://stackoverflow.com/questions/19965795/how-to-write-log-to-file
		logFile, err := os.OpenFile("server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		log.SetOutput(logFile)
		log.Println()
		log.Println("--------------------------------------")
	}

	printProgramName()
	// Log area "startup"
	println("------------------------------------------------------------")
	println("                       Startup")
	println("------------------------------------------------------------")

	// Extract assets from executable
	ExtractAssets()

	dash := NewDashing(root, portStr, os.Getenv("TOKEN")).Start()
	println()
	println("Listen on http://localhost:" + portStr)
	println()
	
	// Log aread "running"
	println("------------------------------------------------------------")
	println("                      Running")
	println("------------------------------------------------------------")

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