package exercices

import (
	"math/rand"
	"time"
	"fmt"
)

func EchoInChan(id int, ch chan<- int) {
	for {
		num := rand.Intn(100)
		fmt.Printf("Goroutine %d écrit %d\n", id, num)
		ch <- num
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	}
}