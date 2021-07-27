package main

import (
	"godashing/internal/utils"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr"
)

// All files in the folder "assets" will be embedded to the binary file while
// compiling because of the usage of "packr.NewBox(...)" in the following lines.
// When this function is executed all embedded files will be extracted an stored
// in the current working directory.
func ExtractAssets() {
	extractAsset(packr.NewBox("../../assets/dashboards"))
	extractAsset(packr.NewBox("../../assets/jobs"))
	extractAsset(packr.NewBox("../../assets/public"))
	extractAsset(packr.NewBox("../../assets/widgets"))
}

// Extract all files in the given box to the current working directory and
// recreate the path given in the box.
// If the base folder (e.g. widgets) already exists, the box will be skipped.
func extractAsset(jobbox packr.Box) {
	// Extract folder name from path
	var assetPath = filepath.Base(jobbox.Path)

	// Skip folder if it already exists
	log.Println("Check for asset folder '" + assetPath + "'")
	if ok, _ := utils.Exists(root + assetPath); ok {
		return
	}

	// Extract every file in the box
	log.Println("Extract asset folder '" + assetPath + "'")
	jobbox.Walk(func(path string, f packr.File) error {		
		// Make all folders
		os.MkdirAll(root + assetPath + string(filepath.Separator) + filepath.Dir(path), 0777)

		// Write data to files
		content, _ := jobbox.FindString(path)
		var filename = root + assetPath + string(filepath.Separator) + path
		if (debugmode) {
			log.Println("Extract asset file '" + filename + "'")
		}
		ioutil.WriteFile(filename, []byte(content), 0644)
		return nil
	})
}