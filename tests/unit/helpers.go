package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"bufio"
	"io/ioutil"
	"net/http"
)

func main() {
	// Get the current working directory
	var cwd string
	err := runtime.Getcwd(&cwd)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the file path of the current file
	filePath := cwd + "/helpers.go"

	// Check if the file exists
	if !strings.HasSuffix(filePath, ".go") {
		fmt.Println("File not found: ", filePath)
		return
	}

	// Read the file contents
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse the file contents
	var helpers []string
	err = strings.Split(string(data), "\n", 0, &helpers)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the helpers
	for _, helper := range helpers {
		fmt.Println(helper)
	}
}