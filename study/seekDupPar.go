// package for searching file duplicates in current or any other directory
package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var (
	root   = flag.String("d", ".", "define directory for searching duplicates")
	del    = flag.Bool("r", false, "remove all duplicates")
	paths  = []string{}
	result = make(map[[32]uint8]string)
	wg     = sync.WaitGroup{}
)

func main() {
	var mu sync.RWMutex
	flag.Parse()
	getFiles(*root)
	wg.Add(len(paths))
	for _, path := range paths {
		go compareFiles(path, &mu, *del)
	}
	wg.Wait()
}

// function GetFiles for getting all files from root directory recursively
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

// function compareFiles function to compare files by sha256summ and print all duplicates except one original file
func compareFiles(path string, mu *sync.RWMutex, del bool) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		wg.Done()
		return
	}
	hash := sha256.Sum256([]byte(data))
	mu.Lock()
	if _, ok := result[hash]; ok {
		fmt.Fprintln(os.Stdout, path)
		if del {
			err := os.Remove(path)
			if err != nil {
				fmt.Fprintf(os.Stdout, "cannot delete file %s due to %v", path, err)
			}
		}
		delete(result, hash)
	}
	result[hash] = path
	mu.Unlock()
	wg.Done()
	return
}
