package main

import "fmt"

func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {
		if i%prime != 0 {
			dst <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go generate(ch)

	for i := 0; i < 100; i++{
		prime := <-ch
		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
