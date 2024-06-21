package exercices



func Sum(numbers []int, ch chan int) {
	total := 0
	for _, number := range numbers {
		total += number
	}
	ch <- total  
}