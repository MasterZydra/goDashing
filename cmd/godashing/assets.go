package main

import (
	"godashing/internal/utils"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr"
)

func ExtractAssets() {
	extractAsset(packr.NewBox("../../assets/dashboards"))
	extractAsset(packr.NewBox("../../assets/jobs"))
	extractAsset(packr.NewBox("../../assets/public"))
	extractAsset(packr.NewBox("../../assets/widgets"))
}

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