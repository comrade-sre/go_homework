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
	root  = flag.String("d", ".", "define directory for searching duplicates")
	del   = flag.Bool("r", false, "remove all duplicates")
	paths = []string{}
	wg    = sync.WaitGroup{}
	ch    = make(chan string, 100)
)

// Struct for storing file hashes and paths
type Set struct {
	sync.RWMutex
	res map[[32]uint8]string
}

func NewSet() *Set {
	return &Set{
		res: make(map[[32]uint8]string),
	}
}
func (s *Set) Add(i [32]uint8, path string) {
	s.Lock()
	s.res[i] = path
	s.Unlock()
}
func (s *Set) Has(i [32]uint8) bool {
	s.Lock()
	_, ok := s.res[i]
	s.Unlock()
	return ok
}
func (s *Set) Del(i [32]uint8) {
	s.Lock()
	delete(s.res, i)
	s.Unlock()
}
func main() {
	flag.Parse()
	var result = NewSet()
	GetFiles(*root)
	wg.Add(len(paths))
	go func() {
		for _, path := range paths {
			ch <- path
		}
		close(ch)
	}()

	for path := range ch {
		go CompareFiles(path, result, *del)
	}
	wg.Wait()

}

// function GetFiles for getting all files from root directory recursively
func GetFiles(root string) (files []os.FileInfo) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	for _, file := range files {
		if file.IsDir() {
			GetFiles(root + "/" + file.Name())
		} else {
			paths = append(paths, root+"/"+file.Name())
		}
	}
	return
}

// function compareFiles function to compare files by sha256summ and print all duplicates except one original file
func CompareFiles(path string, result *Set, del bool) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		wg.Done()
		return
	}
	hash := sha256.Sum256([]byte(data))
	if ok := result.Has(hash); ok {
		fmt.Fprintln(os.Stdout, path)
		if del {
			err := os.Remove(path)
			if err != nil {
				fmt.Fprintf(os.Stdout, "cannot delete file %s due to %v", path, err)
			}
		}
		result.Del(hash)
	}
	result.Add(hash, path)
	wg.Done()
	return
}
