package exercices


import (
		"fmt"
		"math/rand"
		"time"
)

func Echo(s string, i int) {
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	fmt.Printf("Goroutine %d: %s\n", i, s)
}