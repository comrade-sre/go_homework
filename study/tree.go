package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var (
	root  string
	isDir bool
)

func main() {
	if len(os.Args) == 2 {
		root = os.Args[1]
	} else {
		root = "."
	}
	getFiles(root)
}

func getFiles(root string) (files []os.FileInfo) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	for _, file := range files {
		isDir = file.IsDir()
		if isDir {
			fmt.Printf("|_%s\n", file.Name())
			getFiles(root + "/" + file.Name())
		} else {
			fmt.Printf("%8v%s\n", "|_", file.Name())
		}
	}
	return files
}
