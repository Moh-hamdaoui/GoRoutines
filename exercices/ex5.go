package exercices

import (
	"sync"
	"time"
)

var (
	counter int
	mu sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
		time.Sleep(time.Millisecond * 10)
	}
}
