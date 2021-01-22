package sort

import (
	"math/rand"
	"testing"
)

var (
	sortedBig   = []int{15, 47, 59, 81, 81, 89, 106, 162, 211, 237, 258, 274, 287, 300, 318, 387, 408, 425, 445, 456, 466, 495, 511, 528, 540, 541, 694, 728, 790, 831, 847, 887, 888, 947}
	unsortedBig = []int{81, 887, 847, 59, 81, 318, 425, 540, 456, 300, 694, 511, 162, 89, 728, 274, 211, 445, 237, 106, 495, 466, 528, 258, 47, 947, 287, 888, 790, 15, 541, 408, 387, 831}
	choice      = rand.Intn(len(sortedBig))
)

func TestSort(t *testing.T) {
	result := bubbleSort(unsortedBig)
	if result[choice] != sortedBig[choice] {
		t.Errorf("gets %d, wants %d", result[choice], sortedBig[choice])
	}
	result = insertSort(unsortedBig)
	if result[choice] != sortedBig[choice] {
		t.Errorf("gets %d, wants %d", result[choice], sortedBig[choice])
	}

}
func BenchmarkInsert(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < 10; i++ {
		length := rand.Intn(10000)
		s := make([]int, length)
		for j := 0; j < length; j++ {
			s[j] = rand.Intn(1000)
		}
		_ = insertSort(s)
	}

}
func BenchmarkBubble(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < 10; i++ {
		length := rand.Intn(10000)
		s := make([]int, length)
		for j := 0; j < length; j++ {
			s[j] = rand.Intn(1000)
		}
		_ = bubbleSort(s)
	}

}
