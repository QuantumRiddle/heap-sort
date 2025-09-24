package main

import (
	"fmt"
	"math/rand"
)

func main() {



	a := generateArr(20)
	fmt.Println(a)
	a = swap(a)
	fmt.Println(a)

}

func sortTree(arr []int, lenght int) []int {

	var temp int

	for i := lenght / 2; i >= 0; i-- {

		if 2*i+1>lenght {
			continue
		}

		if arr[i] <= arr[2*i+1] {
			temp = arr[i]
			arr[i] = arr[2*i+1]
			arr[2*i+1] = temp
			temp = 0
			
		}

		if 2*i+2 > lenght {
			continue
		}

		if arr[i] <= arr[2*i+2] {
			temp = arr[i]
			arr[i] = arr[2*i+2]
			arr[2*i+2] = temp
			temp = 0
		}

	}

	for i := 0; i < lenght / 2; i++ {

		if arr[i] <= arr[2*i+1] {
			temp = arr[i]
			arr[i] = arr[2*i+1]
			arr[2*i+1] = temp
			temp = 0

		}

		if 2*i+2 > lenght {
			continue
		}

		if arr[i] <= arr[2*i+2] {
			temp = arr[i]
			arr[i] = arr[2*i+2]
			arr[2*i+2] = temp
			temp = 0

		}

	}

	return arr

}

func swap(arr []int) []int {

	var temp int
	l:=len(arr)-1

	for i := l; i >= 0; i-- {
		
		sortTree(arr, l)
		temp = arr[0]
		arr[0] = arr[l]
		arr[l] = temp
		temp = 0
		l--
		

	}

	return arr

}

func generateArr(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}
