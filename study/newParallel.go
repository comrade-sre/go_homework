package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime/trace"
	"sync"
)

var (
	root   = flag.String("d", ".", "define directory for searching duplicates")
	paths  = []string{}
	result = make(map[[32]uint8]string)
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	var mu sync.Mutex
	flag.Parse()
	getFiles(*root)
	for _, path := range paths {
		go compareFiles(path, &mu)

	}
	for key, value := range result {
		fmt.Println("Key:", key, "Value:", value)
	}

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
func compareFiles(path string, mu *sync.Mutex) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	hash := sha256.Sum256([]byte(data))
	mu.Lock()
	defer mu.Unlock()
	if _, ok := result[hash]; ok {
		fmt.Println(path)
		delete(result, hash)
	}
	result[hash] = path
	return
}
