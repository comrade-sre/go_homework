package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	root   = flag.String("d", ".", "define directory for searching duplicates")
	del    = flag.Bool("r", false, "remove all duplicates")
	paths  = []string{}
	result = make(map[[32]uint8]string)
)

func main() {
	flag.Parse()
	getFiles(*root)
	compareFiles(paths, *del)
}
func getFiles(root string) (files []os.FileInfo) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	for _, file := range files {
		if file.IsDir() {
			getFiles(root + "/" + file.Name())
		} else {
			paths = append(paths, root+"/"+file.Name())
		}
	}
	return
}
func compareFiles(s []string, del bool) {
	for _, path := range paths {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		hash := sha256.Sum256([]byte(data))
		if _, ok := result[hash]; ok {
			fmt.Println(path)
			if del {
				err := os.Remove(path)
				if err != nil {
					fmt.Fprintf(os.Stdout, "cannot delete file %s due to %v", path, err)
				}
			}
			delete(result, hash)
		}
		result[hash] = path
	}
}
