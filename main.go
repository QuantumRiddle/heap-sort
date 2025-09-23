package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// a:=generateArr(10)
	// fmt.Println(a)
	a := []int{14, 61, 85, 24, 74, 26, 17, 50, 40, 45, 21, 32, 59, 58, 13}

	a = createTree(a,len(a)-1)


	fmt.Println(a)
}

func generateArr(n int) []int {

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(11)
	}

	return arr

}

func createTree(arr []int, l int) []int {

	var temp int

	lastParrent := l/2 - 1

	for i := lastParrent; i >= 0; i-- {

		temp = 0

		if arr[i] < arr[2*i+1] {

			temp = arr[2*i+1]
			arr[2*i+1] = arr[i]
			arr[i] = temp
			temp = 0

			if arr[i] < arr[2*i+2] {
				temp = arr[2*i+2]
				arr[2*i+2] = arr[i]
				arr[i] = temp
				continue
			}
			continue
		}

	}
	
	
	return arr

}

