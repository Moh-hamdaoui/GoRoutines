package exercices

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


func Producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		num := rand.Intn(100)
		fmt.Printf("Producteur %d produit %d\n", id, num)
		ch <- num
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

func Consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("Consommateur %d traite %d, son carré est %d\n", id, num, num*num)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

