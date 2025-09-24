package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {

	var x string

	for i := 0; i < 5; i++ {
		var a int
		fmt.Println("Введите размер массива:")
		fmt.Fscan(os.Stdin, &a)

		best := best(a)
		_, c, m := swap(best)
		fmt.Printf("Лучший случай\n C:%d M:%d\n", c, m)

		worst := worst(a)
		_, c, m = swap(worst)
		fmt.Printf("Худший случай\n C:%d M:%d\n", c, m)

		rand := generateArr(a)
		_, c, m = swap(rand)
		fmt.Printf("Случайный массив\n C:%d M:%d\n", c, m)
	}

	fmt.Fscan(os.Stdin, &x)

}

func sortTree(arr []int, lenght, compare, move int) ([]int, int, int) {

	var temp int

	for i := lenght / 2; i >= 0; i-- {

		if 2*i+1 > lenght {
			continue
		}

		compare++
		if arr[i] <= arr[2*i+1] {
			temp = arr[i]
			arr[i] = arr[2*i+1]
			arr[2*i+1] = temp
			temp = 0
			move++

		}

		if 2*i+2 > lenght {
			continue
		}

		compare++
		if arr[i] <= arr[2*i+2] {
			temp = arr[i]
			arr[i] = arr[2*i+2]
			arr[2*i+2] = temp
			temp = 0
			move++
		}

	}

	for i := 0; i < lenght/2; i++ {

		compare++
		if arr[i] <= arr[2*i+1] {
			temp = arr[i]
			arr[i] = arr[2*i+1]
			arr[2*i+1] = temp
			temp = 0
			move++
		}

		if 2*i+2 > lenght {
			continue
		}

		compare++
		if arr[i] <= arr[2*i+2] {
			temp = arr[i]
			arr[i] = arr[2*i+2]
			arr[2*i+2] = temp
			temp = 0
			move++

		}

	}

	return arr, compare, move

}

func swap(arr []int) ([]int, int, int) {

	var temp, c, m int
	l := len(arr) - 1

	for i := l; i >= 0; i-- {

		arr, c, m = sortTree(arr, l, c, m)
		temp = arr[0]
		arr[0] = arr[l]
		arr[l] = temp
		temp = 0
		m++
		l--

	}

	return arr, c, m

}

func generateArr(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func worst(n int) []int {

	arr := make([]int, n)
	k := 0

	for i := n; i > 0; i-- {
		arr[k] = i
		k++
	}

	return arr
}

func best(n int) []int {

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
	}

	return arr

}
