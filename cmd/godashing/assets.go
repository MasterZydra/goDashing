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
	var assetPath = filepath.Base(jobbox.Path)

	log.Println("Check for asset folder '" + assetPath + "'")
	if ok, _ := utils.Exists(root + assetPath); ok {
		return
	}

	log.Println("Extract asset folder '" + assetPath + "'")
	jobbox.Walk(func(path string, f packr.File) error {		
		os.MkdirAll(root + assetPath + string(filepath.Separator) + filepath.Dir(path), 0777)

		content, _ := jobbox.FindString(path)
		ioutil.WriteFile(root + assetPath + string(filepath.Separator) + path, []byte(content), 0644)
		return nil
	})
}