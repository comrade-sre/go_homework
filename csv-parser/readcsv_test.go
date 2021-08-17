package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"os"
	"runtime"
	"testing"
	"time"
)

const (
	TestLength  = 1000000
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ\n"
)

var (
	ch = make(chan string, TestLength)
)

func PrintMemUsage() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return fmt.Sprintf("Alloc = %v MiB", m.Alloc/1024/1024)
}
func RandBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		if i%15 == 0 {
			b[i] = letterBytes[len(letterBytes)-1]
			continue
		}
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}
func TestMain(m *testing.M) {
	st := m.Run()
	os.Exit(st)

}
func TestReadCsv(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rand.Seed(time.Now().UnixNano())
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	t.Log(PrintMemUsage())
	data := RandBytes(TestLength)
	PrintMemUsage()
	buf := bytes.NewBuffer(data)
	reader := bufio.NewReader(buf)
	go ReadCsv(*reader, ch, logger, ctx)
	var i int
	for res := range ch {
		i++
		t.Logf("received random line  from channel: %s\n", res)
		t.Log(PrintMemUsage())
	}
	if i == TestLength/15 {
		t.Errorf("not all lines received from channel, wants %d, get %d", TestLength/15, i)
	}
}
