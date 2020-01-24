package main

import (
	"fmt"
	"gitlab.com/stackworx-public/react-static-nginx/pkg"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filename := "/usr/share/nginx/html/index.html"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Printf("%s does not exist", filename)
		os.Exit(-1)
	}

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("Failed to read index.html: %s", err.Error())
	}

	vars, err := pkg.GetEnvVars(os.Environ())

	if err != nil {
		log.Printf("%s", err)
		os.Exit(-1)
	}

	for key, _ := range vars {
		fmt.Printf("Found ENV: %s", key)
	}

	result, err := pkg.ReplaceEnvVars(data, vars)

	if err != nil {
		log.Printf("%s", err)
		os.Exit(-1)
	}

	err = ioutil.WriteFile(filename, result, 0644)

	if err != nil {
		log.Printf("%s", err)
		os.Exit(-1)
	}
}
