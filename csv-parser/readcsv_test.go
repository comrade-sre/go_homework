package main

import (
	"bufio"
	"bytes"
	"go.uber.org/zap"
	"math/rand"
	"testing"
	"time"
)

const (
	TestLength  = 100000
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ\n"
)

var (
	ch = make(chan string, TestLength)
)

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
func TestReadCsv(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	data := RandBytes(TestLength)
	buf := bytes.NewBuffer(data)
	reader := bufio.NewReader(buf)
	go ReadCsv(*reader, ch, logger)
	var i int
	for res := range ch {
		i++
		t.Logf("received random line  from channel: %s\n", res)
	}
	if i == TestLength/15 {
		t.Errorf("not all lines received from channel, wants %d, get %d", TestLength/15, i)
	}
}
