//Assignment 1
//Student Name: Chen, Junhong
//Student Number: 300140321

package main

import (
	"fmt"
	"math/rand"
)
import "sync"

func RandomGenerator(wg *sync.WaitGroup, stop chan bool, m int) chan int {
	intStream := make(chan int)
	go func() {
		defer func() { wg.Done() }()
		defer close(intStream)
		for {
			select {
			case <-stop:
				return
			case intStream <- m * rand.Intn(1000000):
			}
		}
	}()
	return intStream
}

func Multiple(x int, m int) bool {
	if x%m != 0 {
		return false
	}
	return true
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	ch := make(chan bool, 3)
	multiples5 := 0
	multiples13 := 0
	multiples97 := 0
	generator1 := RandomGenerator(&wg, ch, 5)
	generator2 := RandomGenerator(&wg, ch, 13)
	generator3 := RandomGenerator(&wg, ch, 97)

	for i := 0; i < 100; i++ {
		num := 0
		select {
		case num = <-generator1:
			if Multiple(num, 5) {
				multiples5++
			}
		case num = <-generator2:
			if Multiple(num, 13) {
				multiples13++
			}
		case num = <-generator3:
			if Multiple(num, 97) {
				multiples97++
			}
		}

	}

	fmt.Printf("the total number of generated multiples of 5: %d\n", multiples5)
	fmt.Printf("the total number of generated multiples of 13: %d\n", multiples13)
	fmt.Printf("the total number of generated multiples of 97: %d\n", multiples97)

	for j := 0; j < 3; j++ {
		ch <- true
	}
	wg.Wait()
}
