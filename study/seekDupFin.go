// package for searching file duplicates in current or any other directory
package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"go.uber.org/zap"
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
	delete(s.res, i)
}
func main() {
	flag.Parse()
	host,_ := os.Hostname()
	logger, err := NewLogger()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer logger.Sync()
	logger = logger.With(zap.String("hostname: ", host))
	var result = NewSet()
	GetFiles(*root, logger)
	go func() {
		for _, path := range paths {
			ch <- path
		}
		close(ch)
	}()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ch, result, logger)
	}
	wg.Wait()
	paths = nil
	deDup(result, logger)

}
func NewLogger() (*zap.Logger, error) {
  cfg := zap.NewProductionConfig()
  cfg.OutputPaths = []string{
    "/Users/andrei.mironov/test/test.log",
  }
  return cfg.Build()
}

// function for deleting duplicates from map
func deDup(result *Set, logger *zap.Logger) {
	for hash, paths := range result.res {
		length := len(paths)
		if length > 1 {
			for i := 0; i < length-1; i++ {
				if *del {
					err := os.Remove(paths[i])
					if err != nil {
						logger.With(zap.Uint64("uid", 00007)).Error("cannot delete file " + paths[i] + " due to" + err.Error())
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
func worker(ch chan string, result *Set, logger *zap.Logger) {
	for path := range ch {
		CalcHash(path, result, logger)
	}
	wg.Done()
}

// function GetFiles for getting all files from root directory recursively
func GetFiles(root string, logger *zap.Logger) (files []os.FileInfo) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		logger.With(zap.Uint64("uid", 00017)).Error(err.Error())
	}
	for _, file := range files {
		if file.IsDir() {
			GetFiles(root + "/" + file.Name(), logger)
		} else if file.Mode().IsRegular() {
			paths = append(paths, root+file.Name())
			logger.With(zap.Uint64("uid", 00001)).Info("file successfully added " + root+file.Name())
		} else {
			logger.With(zap.Uint64("uid", 00005)).Error("file " + root+file.Name() + " is not a regular file")
		}
	}
	return
}

// function compareFiles function to compare files by sha256summ and add to result map
func CalcHash(path string, result *Set, logger *zap.Logger) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		logger.With(zap.Uint64("uid", 00006)).Error("cannot open " + path)
		return
	} 
	hash := sha256.Sum256([]byte(data))
	result.Add(hash, path)
	logger.With(zap.Uint64("uid", 00002)).Info("hash for " + path + " was added")
	return
}
