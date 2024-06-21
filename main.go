package main

import (
	"GoRoutine/Mohammed/exercices"
	"fmt"
	// "time"
	// // "math/rand"
	// "sync"

)



func main() {

	// Exercice 1
	// var input string
	// fmt.Print("Entrez une chaîne de caractères: ")
	// fmt.Scanln(&input)

	// for i := 1; i <= 5; i++ {
	// 	go exercices.Echo(input, i)
	// }

	// time.Sleep(10 * time.Second) 




	//Exercice 2
	// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// ch := make(chan int)

	// go exercices.Sum(numbers[:len(numbers) / 2], ch)
	// go exercices.Sum(numbers[len(numbers) / 2:], ch)

	// total1 := <-ch
	// total2 := <-ch

	// total := total1 + total2
	// fmt.Println("Somme totale:", total)




	//Exercice 3
	// ch := make(chan int, 5)
	// var wg sync.WaitGroup

	// for i := 1; i <= 3; i++ {
	// 	wg.Add(1)
	// 	go exercices.Producer(i, ch, &wg)
	// }

	// for i := 1; i <= 3; i++ {
	// 	wg.Add(1)
	// 	go exercices.Consumer(i, ch, &wg)
	// }

	// wg.Wait()
	// close(ch)


	// Exercice 4

	ch1 := make(chan int)
	ch2 := make(chan int)


	go exercices.EchoInChan(1, ch1)
	go exercices.EchoInChan(2, ch2)

	
	for {
		select {
		case val := <-ch1:
			fmt.Printf("Reçu du channel 1: %d\n", val)
		case val := <-ch2:
			fmt.Printf("Reçu du channel 2: %d\n", val)
		}
	}
}
