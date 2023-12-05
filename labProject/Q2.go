//Assignment 1
//Student Name: Chen, Junhong
//Student Number: 300140321

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func fct(line []float64) {
	for _, v := range line {
		fmt.Printf("%f, ", v)
	}
}

func fct2(matrix [][]float64) {
	matrix[2][0] = 12345.6
}

func sort(tab []float64) {
	for i := 0; i < len(tab)-1; i++ {
		min := i
		for j := min + 1; j < len(tab); j++ {
			if tab[j] < tab[min] {
				min = j
			}
		}
		tmp := tab[i]
		tab[i] = tab[min]
		tab[min] = tmp
	}
	wg.Done()
}

func transpose(tab [][]float64) {
	length := len(tab)
	tmp := make([][]float64, length)
	for i := 0; i < len(tab); i++ {
		row := make([]float64, length)
		for j := 0; j < len(tab); j++ {
			row[j] = tab[i][j]
		}
		tmp[i] = row
	}
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab); j++ {
			tab[i][j] = tmp[j][i]
		}
	}
}

func sortRows(tab [][]float64) {
	wg.Add(len(tab))
	for _, row := range tab {
		go sort(row)
	}
	wg.Wait()
}

func main() {
	// array := [][]float64{{7.1, 2.3, 1.1},
	//	{4.3, 5.6, 6.8},
	//	{2.3, 2.7, 3.5},
	//	{4.5, 8.1, 6.6}}
	array := [][]float64{{1.1, 7.3, 3.2, 0.3, 3.1},
		{4.3, 5.6, 1.8, 5.3, 3.1},
		{1.3, 2.7, 3.5, 9.3, 1.1},
		{7.5, 5.1, 0.6, 2.3, 3.9}}

	for _, line := range array {
		fct(line)
		fmt.Println()
	}

	sortRows(array)
	transpose(array)
	sortRows(array)
	transpose(array)

	for _, line := range array {
		fct(line)
		fmt.Println()
	}

}
