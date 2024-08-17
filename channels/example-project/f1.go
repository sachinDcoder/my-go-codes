package main

import "fmt"

func main() {
	ch := make(chan int)
	result := make(chan int)

	go producer(ch, 10)
	go consumer(ch, result)

	fmt.Print(<-result)

}

func producer(ch chan<- int, num int) {
	for i := 1; i <= num; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int, result chan<- int) {
	sum := 0
	for num := range ch {
		sum += num
	}
	result <- sum
	close(result)
}
