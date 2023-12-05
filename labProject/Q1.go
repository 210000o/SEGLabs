//Assignment 1
//Student Name: Chen, Junhong
//Student Number: 300140321

package main

import (
	"fmt"
	"math"
	"time"
)

type Point struct {
	x float64
	y float64
}

func MidPoint(p1 Point, p2 Point, ch chan bool) {
	midX := (p1.x + p2.x) / 2
	midY := (p1.y + p2.y) / 2
	length := math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
	fmt.Println("Mid-point: (", midX, ",", midY, ")")
	fmt.Printf("Length: %.2f\n", length)
	ch <- true
}

func main() {
	points := []Point{{8., 1.}, {3., 2.}, {7., 4.}, {6., 3.}}

	ch := make(chan bool)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			fmt.Printf("point= %v %v \n", points[i], points[j])
			go MidPoint(points[i], points[j], ch)
		}
	}
	for i := 0; i < 6; i++ {
		time.Sleep(100 * time.Millisecond)
		<-ch
	}

	fmt.Println("All threads are completed")
}
