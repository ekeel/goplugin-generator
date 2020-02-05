package main

import (
	"flag"
	"fmt"
	"goplugin-generator/template"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	namePtr := flag.String("name", "plugin", "The name of the plugin.")
	directoryPtr := flag.String("dir", "./", "The directory to output the generated plugin to.")

	flag.Parse()

	if _, err := os.Stat(*directoryPtr); os.IsNotExist(err) {
		os.MkdirAll(*directoryPtr, os.ModePerm)
	}

	filePath := path.Join(*directoryPtr, fmt.Sprintf("%s.go", *namePtr))

	err := ioutil.WriteFile(filePath, []byte(template.PluginTemplate), 0755)
	if err != nil {
		panic(err)
	}

	log.Println(fmt.Sprintf("Created plugin at: %s", filePath))
}
