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
	res map[[32]uint8][]string
}

func NewSet() *Set {
	return &Set{
		res: make(map[[32]uint8][]string),
	}
}
func (s *Set) Add(i [32]uint8, path string) {
	s.Lock()
	s.res[i] = append(s.res[i], path)
	s.Unlock()
}
func (s *Set) Has(i [32]uint8) bool {
	s.Lock()
	_, ok := s.res[i]
	s.Unlock()
	return ok
}
func (s *Set) Del(i [32]uint8) {
	//	s.Lock()
	delete(s.res, i)
	//	s.Unlock()
}
func main() {
	flag.Parse()
	var result = NewSet()
	GetFiles(*root)
	go func() {
		for _, path := range paths {
			ch <- path
		}
		close(ch)
	}()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ch, result)
	}
	wg.Wait()
	deDup(result)

}

// function for deleting duplicates from map
func deDup(result *Set) {
	for hash, paths := range result.res {
		length := len(paths)
		if length > 1 {
			for i := 0; i < length-1; i++ {
				if *del {
					err := os.Remove(paths[i])
					if err != nil {
						fmt.Fprintf(os.Stdout, "cannot delete file %s due to %v\n", paths[i], err)
					}
					fmt.Fprintf(os.Stdout, "%s deleted\n", paths[i])
				} else {
					fmt.Fprintln(os.Stdout, paths[i])
				}
			}
		}
		result.Del(hash)
	}
}
func worker(ch chan string, result *Set) {
	for path := range ch {
		CompareFiles(path, result)
	}
	wg.Done()
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
		} else if file.Mode().IsRegular() {
			paths = append(paths, root+"/"+file.Name())
		} else {
			fmt.Fprintf(os.Stderr, "%s is not a regular file\n", root+"/"+file.Name())
		}
	}
	return
}

// function compareFiles function to compare files by sha256summ and add to result map
func CompareFiles(path string, result *Set) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	hash := sha256.Sum256([]byte(data))
	result.Add(hash, path)
	return
}
