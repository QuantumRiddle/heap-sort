package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// a:=generateArr(10)
	// fmt.Println(a)
	//a := []int{14, 61, 85, 24, 74, 26, 17, 50, 40, 45, 21, 32, 59, 58, 13}  [16 13 1 13 10 4] [6 1 1 13 6 16] 
	//						4
	//				8				2
	//			2		4		6		3
	//		9	   2  10 10   9   6   1	  6
	//
	a := []int{6, 1, 1, 13, 6, 16}

	// for i := 0; i < 5; i++ {

	// a := generateArr(6)
	fmt.Println("unsorted ", a)

	a = createTree(a, len(a)-1)
	for i := 0; i < len(a)-2; i++ {
		last := len(a) - 1 - i
		temp := 0
		temp = a[0]
		a[0] = a[last]
		a[last] = temp
		last--
		a = brush(a, last)
	}

	fmt.Println("sorted: ", a)

	// }

}

func generateArr(n int) []int {

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(20)
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
			}
			
		}
		if arr[i] < arr[2*i+2] {
			temp = arr[2*i+2]
			arr[2*i+2] = arr[i]
			arr[i] = temp

			if arr[i] < arr[2*i+1] {
				temp = arr[2*i+1]
				arr[2*i+1] = arr[i]
				arr[i] = temp
				
			}
			
		}
		

	}

	for i := 0; i < lastParrent; i++ {
		temp = 0

		if arr[i] < arr[2*i+1] {

			temp = arr[2*i+1]
			arr[2*i+1] = arr[i]
			arr[i] = temp
			temp = 0
			// if i > 1 {
			// 	i -= 2
			// }

			if arr[i] < arr[2*i+2] {
				temp = arr[2*i+2]
				arr[2*i+2] = arr[i]
				arr[i] = temp
				// if i > 1 {
				// 	i -= 2
				// }
				
			}
			
		}
		if arr[i] < arr[2*i+2] {
			temp = arr[2*i+2]
			arr[2*i+2] = arr[i]
			arr[i] = temp
			// if i > 1 {
			// 	i -= 2
			// }
			if arr[i] < arr[2*i+1] {
				temp = arr[2*i+1]
				arr[2*i+1] = arr[i]
				arr[i] = temp
				// if i > 1 {
				// 	i -= 2
				// }
			
			}
			
		}
		
	}

	return arr

}

func brush(arr []int, l int) []int {

	var temp int

	lastParrent := l / 2

	for i := 0; i < lastParrent; i++ {

		if arr[i] < arr[2*i+1] {

			temp = arr[2*i+1]
			arr[2*i+1] = arr[i]
			arr[i] = temp
			temp = 0
			// if i > 1 {
			// 	i -= 2
			// }

			if arr[i] < arr[2*i+2] {
				temp = arr[2*i+2]
				arr[2*i+2] = arr[i]
				arr[i] = temp
				// if i > 1 {
				// 	i -= 2
				// }
		
			}
		
		}

		if arr[i] < arr[2*i+2] {
			temp = arr[2*i+2]
			arr[2*i+2] = arr[i]
			arr[i] = temp
			// if i > 1 {
			// 	i -= 2
			// }

			if arr[i] < arr[2*i+1] {
				temp = arr[2*i+1]
				arr[2*i+1] = arr[i]
				arr[i] = temp
				// if i > 1 {
				// 	i -= 2
				// }
		
			}
		
		}
		
	}

	createTree(arr,l)

	return arr
}
