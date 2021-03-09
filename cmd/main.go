package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	dashing "goDashing"
	_ "goDashing/jobs"
)

func main() {
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

	dash := dashing.NewDashing(webroot, port, os.Getenv("TOKEN")).Start()
	log.Println("listening on :" + port)

	http.Handle("/", dash)

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
