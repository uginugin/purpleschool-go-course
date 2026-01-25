package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func main() {
	count := 10
	ch := make(chan int, count)
	squaredCh := make(chan int, count)

	go send(count, ch)

	go recv(count, ch, squaredCh)

	for i := 0; i < count; i++ {
		fmt.Print(<-squaredCh, ' ')
	}
}

func send(count int, ch chan int) {
	slice := make([]int, count)

	for i := 0; i < count; i++ {
		slice[i] = rand.IntN(100)
	}

	for i := 0; i < count; i++ {
		ch <- slice[i]
	}
}

func recv(count int, ch chan int, mainCh chan int) {

	for i := 0; i < count; i++ {
		digit := <-ch
		mainCh <- int(math.Pow(float64(digit), 2))
	}
}
